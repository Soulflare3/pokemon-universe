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
	"fmt"
	"net"
	punet "network"
)

const (
	//this should be more than enough
	MESSAGE_BUFFERSIZE = 1000
)

type IProtocol interface {
	ProcessPacket(_packet *punet.Packet)
	ProcessMessage(_message *punet.Message)
	SendLogin(_username string, _password string)
}

type PU_Connection struct {
	socket      net.Conn
	msgChan     chan *punet.Message
	tranceiver  *punet.Tranceiver
	protocol    IProtocol
	loginStatus int
	connected   bool
}

func NewConnection() *PU_Connection {
	return &PU_Connection{msgChan: make(chan *punet.Message, MESSAGE_BUFFERSIZE),
		protocol:    NewGameProtocol(),
		loginStatus: LOGINSTATUS_IDLE}
}

func (c *PU_Connection) Connect() bool {
	//TODO: read from config file
	//ip := "94.75.231.83" //arceus
	ip := "127.0.0.1"
	port := "1337"

	var err error
	c.socket, err = net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Printf("Connection error: %v\n", err.Error())
		return false
	}

	c.connected = true
	c.tranceiver = punet.NewTranceiver(c.socket)
	go c.ReceiveMessages()

	return true
}

func (c *PU_Connection) Close() {
	if c.socket != nil {
		c.socket.Close()
	}
	c.connected = false
}

func (c *PU_Connection) ReceiveMessages() {
	for c.connected {
		if message, err := c.tranceiver.Receive(); err != "" {
			fmt.Printf("Error receiving message: %s\n", err)
			break
		} else {
			select {
			case c.msgChan <- message:
				//great success

			default:
				fmt.Printf("Error: Message buffer full\n")
			}
		}
	}
}

func (c *PU_Connection) HandleMessage() {
	if !c.connected {
		return
	}

	//check if there's a message in the buffer and process it
	//repeat until buffer is empty
	for {
		var breakloop bool
		select {
		case message := <-c.msgChan:
			c.protocol.ProcessMessage(message)

		default:
			breakloop = true
		}
		if breakloop {
			break
		}
	}
}

func (c *PU_Connection) SendMessage(_message *punet.Message) {
	c.tranceiver.Send(_message)
}

func (c *PU_Connection) Game() *PU_GameProtocol {
	protocol, _ := c.protocol.(*PU_GameProtocol)
	return protocol
}
