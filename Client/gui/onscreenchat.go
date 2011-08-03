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
	list "container/vector"
)

type PU_OnscreenChat struct {
	PU_GuiElement

	lastTicks uint32
	messages  list.Vector
}

func NewOnscreenChat() *PU_OnscreenChat {
	chat := &PU_OnscreenChat{}
	chat.visible = true
	g_gui.AddElement(chat)

	chat.lastTicks = sdl.GetTicks()

	return chat
}

func (g *PU_OnscreenChat) Draw() {
	if !g.visible {
		return
	}

	if g_game.self == nil {
		return
	}

	ticks := int(sdl.GetTicks() - g.lastTicks)
	g.lastTicks = sdl.GetTicks()

	for i := 0; i < g.messages.Len(); {
		msg, ok := g.messages.At(i).(*PU_OnscreenChatMessage)
		if ok {
			if !msg.Draw(ticks) {
				g.messages.Delete(i)
			} else {
				i++
			}
		}
	}
}

func (g *PU_OnscreenChat) Add(_name string, _text string) {
	sender := g_map.GetPlayerByName(_name)
	if sender == nil {
		return
	}

	if g == nil {
		return
	}

	for i := 0; i < g.messages.Len(); i++ {
		msg, ok := g.messages.At(i).(*PU_OnscreenChatMessage)
		if ok {
			if msg.name == _name && msg.x == int(sender.GetX()) && msg.y == int(sender.GetY()) {
				msg.AddText(_text)
				return
			}
		}
	}

	g.messages.Push(NewOnscreenChatMessageExt(_name, int(sender.GetX()), int(sender.GetY()), _text))
}

func (g *PU_OnscreenChat) MouseDown(_x int, _y int) {

}

func (g *PU_OnscreenChat) MouseUp(_x int, _y int) {

}

func (g *PU_OnscreenChat) MouseMove(_x int, _y int) {

}

func (g *PU_OnscreenChat) MouseScroll(_dir int) {

}

func (g *PU_OnscreenChat) Focusable() bool {
	return false
}

func (g *PU_OnscreenChat) KeyDown(_keysym int, _scancode int) {

}
