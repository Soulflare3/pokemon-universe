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
)

type PU_OnscreenChat struct {
	PU_GuiElement

	lastTicks uint32
	messages  []*PU_OnscreenChatMessage
}

func NewOnscreenChat() *PU_OnscreenChat {
	chat := &PU_OnscreenChat{}
	chat.visible = true
	chat.messages = make([]*PU_OnscreenChatMessage, 0)
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

	for i := 0; i < len(g.messages); {
		msg := g.messages[i]
		if msg != nil {
			if !msg.Draw(ticks) {
				g.messages = append(g.messages[:i], g.messages[i+1:]...)
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

	for i := 0; i < len(g.messages); i++ {
		msg := g.messages[i]
		if msg != nil {
			if msg.name == _name && msg.x == int(sender.GetX()) && msg.y == int(sender.GetY()) {
				msg.AddText(_text)
				return
			}
		}
	}
	
	g.messages = append(g.messages, NewOnscreenChatMessageExt(_name, int(sender.GetX()), int(sender.GetY()), _text))
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
