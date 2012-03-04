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
package pubattle

import (
	"fmt"

	"io"
	"net"
	pnet "network"
)

type POClientSocket struct {
	socket     net.Conn
	owner      *POClient
	connected  bool
	packetChan chan *pnet.QTPacket
}

func NewPOClientSocket(_owner *POClient) *POClientSocket {
	return &POClientSocket{owner: _owner,
		packetChan: make(chan *pnet.QTPacket, 1000)}
}

func (s *POClientSocket) Connect(_inIpAddr string, _inPortNum string) bool {
	var err error
	s.socket, err = net.Dial("tcp", _inIpAddr+":"+_inPortNum)
	if err != nil {
		fmt.Println("[WARNING] Could not connect to PO battle server")
		fmt.Printf("%v\n\r", err)
		return false
	}

	s.connected = true
	go s.ReceiveMessages()

	loginPacket := s.owner.meLoginPlayer.WritePacket()
	s.SendMessage(loginPacket, COMMAND_Login)

	//s.loginTest()

	return true
}

func (s *POClientSocket) loginTest() {
	packet := pnet.NewQTPacket()
	packet.AddString("HerpDerp")    // Name
	packet.AddString("Dark Info")   // Info
	packet.AddString("Dark Lose")   // Lose text
	packet.AddString("Dark Winrar") // Win text
	packet.AddUint16(0)             // Avatar
	packet.AddString("1")           // Default Tier
	packet.AddUint8(5)              // Generation

	// TEAM - Loop Pokemon
	packet.AddUint16(16)       // pokemon number
	packet.AddUint8(0)         // sub number (alt-forms)
	packet.AddString("Pidgey") // nickname
	packet.AddUint16(0)        // item
	packet.AddUint16(65)       // ability
	packet.AddUint8(0)         // nature
	packet.AddUint8(1)         // gender
	packet.AddUint8(0)         // shiny	
	packet.AddUint8(0)         // happiness
	packet.AddUint8(100)       // level

	// Team - Loop Pokemon - Loop Moves
	packet.AddUint32(16) // moveid
	packet.AddUint32(0)  // moveid
	packet.AddUint32(0)  // moveid
	packet.AddUint32(0)  // moveid
	// Team - Loop Pokemon - End Loop Moves
	// Team - Loop Pokemon - Loop DV
	packet.AddUint8(15) // hp
	packet.AddUint8(15) // attack
	packet.AddUint8(15) // defense
	packet.AddUint8(15) // SpAttack
	packet.AddUint8(15) // SpDefence
	packet.AddUint8(15) // Speed
	// Team - Loop Pokemon - End Loop DV
	// Team - Loop Pokemon - Loop EV
	packet.AddUint8(152)
	packet.AddUint8(92)
	packet.AddUint8(60)
	packet.AddUint8(64)
	packet.AddUint8(64)
	packet.AddUint8(76)
	// Team - Loop Pokemon - End Loop EV

	// Loop 5 more empty pokemon
	for i := 0; i < 5; i++ {
		packet.AddUint16(0)  // pokemon number
		packet.AddUint8(0)   // sub number (alt-forms)
		packet.AddString("") // nickname
		packet.AddUint16(0)  // item
		packet.AddUint16(0)  // ability
		packet.AddUint8(0)   // nature
		packet.AddUint8(0)   // gender
		packet.AddUint8(0)   // shiny
		packet.AddUint8(0)   // happiness
		packet.AddUint8(0)   // level	

		packet.AddUint32(0) // moveid
		packet.AddUint32(0) // moveid
		packet.AddUint32(0) // moveid
		packet.AddUint32(0) // moveid

		packet.AddUint8(0) // hp
		packet.AddUint8(0) // attack
		packet.AddUint8(0) // defense
		packet.AddUint8(0) // SpAttack
		packet.AddUint8(0) // SpDefence
		packet.AddUint8(0) // Speed

		packet.AddUint8(0)
		packet.AddUint8(0)
		packet.AddUint8(0)
		packet.AddUint8(0)
		packet.AddUint8(0)
		packet.AddUint8(0)
	}

	// Team - End Loop Pokemon
	packet.AddUint8(1)  // Ladder
	packet.AddUint8(1)  // Show team
	packet.AddUint32(1) // Colour

	s.SendMessage(packet, COMMAND_Login)
}

func (s *POClientSocket) Disconnect() {
	// Update pokemon data
	s.owner.UpdatePokemonData()
	
	// Close network connection
	s.connected = false
	s.socket.Close()
	close(s.packetChan)
}

func (s *POClientSocket) SendMessage(_buffer pnet.IPacket, _header int) {
	packet := pnet.NewQTPacket()
	packet.AddUint8(uint8(_header))
	if !packet.AddBuffer(_buffer.GetBufferSlice()) {
		fmt.Println("[ERROR} PACKET IS TOO LARGE, CAN NOT ADD BUFFER!")
		return
	}
	packet.SetHeader()

	// Send message to the big bad internetz and pray for it to arrive
	s.socket.Write(packet.Buffer[0:packet.MsgSize])
}

func (s *POClientSocket) ReceiveMessages() {
	for s.connected {
		var headerbuffer [2]uint8
		recv, err := s.socket.Read(headerbuffer[0:])
		if err != nil || recv == 0 {
			fmt.Println("[POCLIENTSOCKET] Disconnected")
			break
		}

		packet := pnet.NewQTPacket()
		copy(packet.Buffer[0:2], headerbuffer[0:2])
		packet.GetHeader()

		databuffer := make([]uint8, packet.MsgSize)

		reloop := false
		bytesReceived := uint16(0)
		for bytesReceived < packet.MsgSize {
			recv, err := io.ReadFull(s.socket, databuffer[bytesReceived:])
			if recv == 0 {
				reloop = true
				break
			} else if err != nil {
				fmt.Printf("[POCLIENTSOCKET] Read error: %v\n", err)
				reloop = true
				break
			}
			bytesReceived += uint16(recv)
		}

		if reloop {
			continue
		}
		copy(packet.Buffer[2:], databuffer[:])

		s.owner.ProcessPacket(packet)
	}

	fmt.Println("EXIT")

	s.connected = false
}
