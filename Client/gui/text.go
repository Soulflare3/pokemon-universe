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

//The PU_Text type is a piece of text consisting of different parts that can have different colors
//This is mainly for use in the chatbox

type PU_Textpart struct {
	text string
	color uint32
}

func NewTextpart(_text string, _color uint32) *PU_Textpart {
	return &PU_Textpart{text : _text, color : _color}
}

type PU_Text struct {
	font *PU_Font
	parts map[int]*PU_Textpart
	count int
}

func NewText(_font int) *PU_Text{
	text := &PU_Text{}
	text.font = g_engine.GetFont(_font)
	text.parts = make(map[int]*PU_Textpart)
	return text
} 

func (t *PU_Text) Add(_text string, _color uint32) {
	t.parts[t.count] = NewTextpart(_text, _color)
	t.count++
}

func (t *PU_Text) AddToLast(_text string) {
	part, present := t.parts[t.count-1]
	if present {
		part.text += _text
	}
}

func (t *PU_Text) GetPart(_index int) *PU_Textpart {
	part, present := t.parts[_index]
	if present {
		return part
	}
	return nil
}

func (t *PU_Text) GetAll() string {
	str := ""
	for i := 0; i < t.count; i++ {
		part, present := t.parts[i]
		if present {
			str += part.text
		}
	}
	return str
}

