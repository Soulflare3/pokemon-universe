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
	"sdl"
	list "container/list"
)

type PU_Listbox struct {
	PU_GuiElement

	font       *PU_Font
	color      sdl.Color
	bgcolor    sdl.Color
	itemcolor  sdl.Color
	bold       bool
	italic     bool
	underlined bool

	items *list.List

	ItemSelectedCallback func(_item int)
	KeyDownCallback      func(_keysym int, _scancode int)
}

func NewListbox(_rect *PU_Rect, _font int) *PU_Listbox {
	listbox := &PU_Listbox{font: g_engine.GetFont(_font),
		items: list.New()}
	listbox.rect = _rect
	listbox.visible = true
	listbox.SetColor(255, 255, 255)
	listbox.SetBgColor(0, 0, 0)
	listbox.SetItemColor(95, 95, 95)
	g_gui.AddElement(listbox)
	return listbox
}

func (l *PU_Listbox) SetFont(_id int) {
	l.font = g_engine.GetFont(_id)
}

func (l *PU_Listbox) SetColor(_red uint8, _green uint8, _blue uint8) {
	l.color.R = _red
	l.color.G = _green
	l.color.B = _blue
}

func (l *PU_Listbox) SetBgColor(_red uint8, _green uint8, _blue uint8) {
	l.bgcolor.R = _red
	l.bgcolor.G = _green
	l.bgcolor.B = _blue
}

func (l *PU_Listbox) SetItemColor(_red uint8, _green uint8, _blue uint8) {
	l.itemcolor.R = _red
	l.itemcolor.G = _green
	l.itemcolor.B = _blue
}

func (l *PU_Listbox) SetStyle(_bold bool, _italic bool, _underlined bool) {
	l.bold, l.italic, l.underlined = _bold, _italic, _underlined
}

func (l *PU_Listbox) Draw() {
	if !l.visible {
		return
	}

	//draw the background
	g_engine.DrawFillRect(l.rect, &l.bgcolor, 200)

}

func (l *PU_Listbox) MouseDown(_x int, _y int) {

}

func (l *PU_Listbox) MouseUp(_x int, _y int) {
	if l.rect.Contains(_x, _y) {
		g_gui.SetFocus(l)
	}
}

func (l *PU_Listbox) MouseMove(_x int, _y int) {

}

func (l *PU_Listbox) MouseScroll(_dir int) {

}

func (l *PU_Listbox) Focusable() bool {
	return true
}

func (l *PU_Listbox) KeyDown(_keysym int, _scancode int) {
	if !l.focus {
		return
	}

	if l.KeyDownCallback != nil {
		l.KeyDownCallback(_keysym, _scancode)
	}
}
