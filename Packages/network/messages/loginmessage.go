package network

import "os"

type LoginMessage struct {
	// Receive
	Username 		string
	Password 		string
	ClientVersion 	uint16
	
	// Send
	Status			uint32
}

// GetHeader returns the header value of this message
func (m *LoginMessage) GetHeader() uint8 {
	return HEADER_LOGIN
}

// ReadPacket reads all data from a packet and puts it in the object
func (m *LoginMessage) ReadPacket(_packet *Packet) os.Error {
	m.Username = _packet.ReadString()
	m.Password = _packet.ReadString()
	m.ClientVersion = _packet.ReadUint16()
	
	return nil
}

// WritePacket write the needed object data to a Packet and returns it
func (m *LoginMessage) WritePacket() (*Packet, os.Error) {
	packet := NewPacketExt(HEADER_LOGIN)
	packet.AddUint32(m.Status)
	
	return packet, nil
}
