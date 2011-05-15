package main

import "fmt"

type MoveMessagesPart []string

type MoveInfo struct {
	Names	map[uint16]string
	MoveMessages map[uint16]MoveMessagesPart
}

func NewMoveInfo() *MoveInfo {
	info := &MoveInfo{ Names: make(map[uint16]string),
				MoveMessages: make(map[uint16]MoveMessagesPart) }
	info.init()
	return info
}

func (m *MoveInfo) init() {
	m.Names[16] = "Gust"
	m.Names[338] = "Frenzy Plant"
}

func (m *MoveInfo) GetMoveName(_moveNumber uint16) string {
	value, found := m.Names[_moveNumber]
	
	if !found {
		fmt.Printf("ERROR - Could not find move: %d\n", _moveNumber)
		return "Unknown Move"
	}
	
	return value
}

func (m *MoveInfo) GetMoveMessage(_move uint16, _part uint8) string {
	value, found := m.MoveMessages[_move]
	
	if !found {
		fmt.Printf("ERROR - Could not find move message %d part %d\n", _move, _part)
	}
	
	return value[_part]
}