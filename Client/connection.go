/*Pokemon Universe MMORPG
Copyright (C) 2010 the Pokemon Universe Authors

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.*/
package main

import (
	"os"
	"fmt"
	"net"
	punet "network"
	"io"
)

const (
	//this should be more than enough
	PACKET_BUFFERSIZE = 1000
)

type IProtocol interface {
	ProcessPacket(_packet *punet.Packet)
	SendLogin(_username string, _password string)
}

type PU_Connection struct {
	socket net.Conn
	packetChan chan *punet.Packet
	protocol IProtocol
	loginStatus int
	connected bool
}

func NewConnection() *PU_Connection {
	return &PU_Connection{packetChan : make(chan *punet.Packet, PACKET_BUFFERSIZE),
						  protocol : NewGameProtocol(),
						  loginStatus : LOGINSTATUS_IDLE}
}

func (c *PU_Connection) Connect() bool {
	//TODO: read from config file
	ip := "94.75.231.83" //arceus
	//port := "6161"
	port := "6666"
	
	var err os.Error
	c.socket, err = net.Dial("tcp", "", ip+":"+port)
	if err != nil {
		fmt.Printf("Connection error: %v\n", err.String())
		return false
	}
	
	c.connected = true
	go c.ReceivePackets()
	
	return true
}

func (c *PU_Connection) Close() {
	if c.socket != nil {
		c.socket.Close()
	}
	c.connected = false
}

func (c *PU_Connection) ReceivePackets() {
	for c.connected {
		var headerbuffer [2]uint8 
		recv, err := io.ReadFull(c.socket, headerbuffer[0:])
		if err != nil || recv == 0 {
			fmt.Printf("Disconnected\n")
			break
		}

		packet := punet.NewPacket()
		copy(packet.Buffer[0:2], headerbuffer[0:2])
		packet.GetHeader()
		
		databuffer := make([]uint8, packet.MsgSize)
		
		reloop := false
		bytesReceived := uint16(0)
		for bytesReceived < packet.MsgSize {
			recv, err = io.ReadFull(c.socket, databuffer[bytesReceived:])
			if recv == 0 {	
				reloop = true
				break 
			} else if err != nil {
				fmt.Printf("Connection read error: %v\n", err)
				reloop = true
				break
			}
			bytesReceived += uint16(recv)
		}
		if reloop {
			continue
		}
		
		copy(packet.Buffer[2:], databuffer[:])
		
		//put the packet in the buffer
		select {
			case c.packetChan <- packet:
				//great success
				
			default:
				fmt.Printf("Error: Packet buffer full\n")
		}
	}
}

func (c *PU_Connection) HandlePacket() {
	if !c.connected {
		return
	}
	
	//check if there's a packet in the buffer and process it
	//repeat until buffer is empty
	for {
		var breakloop bool
		select {
			case packet := <- c.packetChan:
				c.protocol.ProcessPacket(packet)
				
			default:
				breakloop = true
		}
		if breakloop {
			break
		}
	}
}

func (c *PU_Connection) SendMessage(_message punet.INetMessageWriter) {
	packet, _ := _message.WritePacket()
	packet.SetHeader()
	
	c.socket.Write(packet.Buffer[0:packet.MsgSize])
}

func (c *PU_Connection) Game() *PU_GameProtocol {
	protocol, _ := c.protocol.(*PU_GameProtocol)
	return protocol
}
