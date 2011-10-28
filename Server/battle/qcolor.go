package main

import (
	pnet "network"
)

type QColor struct {
	Spec uint8
	Alpha uint16
	Red uint16
	Green uint16
	Blue uint16
	Pad uint16
	Html string
}

func NewQColor() *QColor {
	return &QColor { Spec: 0, Alpha: uint16(0xFFFF), 
					 Red: 0, Green: 0, Blue: 0,
					 Pad: 0, Html: ">" }
}

func NewQColorFromPacket(_packet *pnet.QTPacket) *QColor {
	color := QColor {}
	color.Spec = _packet.ReadUint8()
	color.Alpha = _packet.ReadUint16()
	color.Red = _packet.ReadUint16()
	color.Green = _packet.ReadUint16()
	color.Blue = _packet.ReadUint16()
	color.Pad = _packet.ReadUint16()
	color.Html = ">"
}

func (c *QColor) WritePacket() pnet.IPacket {
	packet := pnet.NewQTPacket()
	packet.AddUint8(c.Spec)
	packet.AddUint16(c.Alpha)
	packet.AddUint16(c.Red)
	packet.AddUint16(c.Green)
	packet.AddUint16(c.Blue)
	packet.AddUint16(c.Pad)
	return packet
}