package main

import (
	"net"
	"os"
	"io"
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
	s.socket, err = net.Dial("tcp", _inIpAddr + ":" + _inPortNum)
	if err != nil {
		g_logger.Println("[WARNING] Could not connect to PO battle server")
		return false
	}
	
	s.connected = true
	go s.ReceiveMessages()
	
	return true
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
			recv, err := io.ReadFull(c.socket, databuffer[bytesReceived:])
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
			case packet := <- s.packetChan
				s.owner.ProcessPacket(packet)
			
			default:
				breakloop = true
		}
		if breakloop {
			break
		}
	}
}