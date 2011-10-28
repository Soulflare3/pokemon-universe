package main

import (
	"net"
	"os"
	"io"
	"fmt"
	pnet "network"
)

type POClientSocket struct {
	socket		net.Conn
	owner		*POClient
	connected	bool
	packetChan	chan *pnet.QTPacket
}

func NewPOClientSocket(_owner *POClient) *POClientSocket {
	return &POClientSocket { owner: _owner,
							 packetChan: make(chan *pnet.QTPacket, 1000) }
}

func (s *POClientSocket) Connect(_inIpAddr string, _inPortNum int) bool {
	var err os.Error
	s.socket, err = net.Dial("tcp", _inIpAddr + ":" + string(_inPortNum))
	if err != nil {
		fmt.Println("[WARNING] Could not connect to PO battle server")
		return false
	}
	
	s.connected = true
	go s.ReceiveMessages()
	
	loginPacket := s.owner.meLoginPlayer.WritePacket()
	s.SendMessage(loginPacket, COMMAND_Login)
	
	return true
}

func (s *POClientSocket) Disconnect() {
	//
	// TOOD: Send Logout message
	//
	
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
		var headerbuffer [1]uint8
		recv, err := io.ReadFull(s.socket, headerbuffer[0:])
		if err != nil || recv == 0 {
			fmt.Println("[POCLIENTSOCKET] Disconnected")
			break
		}
		
		packet := pnet.NewQTPacket()
		copy(packet.Buffer[0:1], headerbuffer[0:1])
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
		
		// Put the packet in the buffer
		select {
			case s.packetChan <- packet:
				// done
			default:
				fmt.Println("[POCLIENTSOCKET] ERROR: Packet buffer full!")
		}
	}
	
	s.connected = false
}

func (s *POClientSocket) HandlePacket() {
	if !s.connected {
		return
	}
	
	for {
		var breakloop bool
		select {
			case packet := <- s.packetChan:
				s.owner.ProcessPacket(packet)
			
			default:
				breakloop = true
		}
		if breakloop {
			break
		}
	}
}