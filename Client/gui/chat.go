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

const (
	SPEAK_NORMAL = 1
	SPEAK_YELL = 2
	SPEAK_WHISPER = 3
	SPEAK_PRIVATE = 6
)

type PU_Chat struct {
	PU_GuiElement
	
	activeChannel int 
	
	channels map[int]*PU_ChatChannel
	
	listVisible bool
	listBlinkLast uint32
	listBlink bool
}

func NewChat() *PU_Chat {
	chat := &PU_Chat{channels : make(map[int]*PU_ChatChannel),
					 listBlinkLast : sdl.GetTicks()}
	chat.visible = true
	g_gui.AddElement(chat)
	return chat
}

func (c *PU_Chat) SetActive(_id int) {
	for id, channel := range c.channels {
		if id == _id {
			channel.SetActive(true)
			channel.updated = false
		} else {
			channel.SetActive(false)
		}
	}
	c.activeChannel = _id
}

func (c *PU_Chat) AddChannel(_id int, _name string, _closable bool) {
	if _, exists := c.channels[_id]; !exists {
		channel := NewChatChannel(_id, _name)
		channel.closable = _closable
		c.channels[_id] = channel
	}
}

func (c *PU_Chat) CloseChannel(_id int) {
	if channel, exists := c.channels[_id]; exists {
		if channel.id == c.activeChannel {
			c.activeChannel = 0
		}
		
		channel.Close()
		c.channels[_id] = channel, false
	}
}

func (c *PU_Chat) AddMessage(_id int, _text *PU_Text) {
	if channel, exists := c.channels[_id]; exists {
		channel.AddMessage(_text)
		if channel.id != c.activeChannel && channel.notifications {
			channel.updated = true
		}
	}
}

func (c *PU_Chat) DrawChatList() {
	if sdl.GetTicks()-c.listBlinkLast >= 500 {
		c.listBlink = !c.listBlink
		c.listBlinkLast = sdl.GetTicks()
	}
	
	hasUpdate := false
	if c.listVisible {
		//top of the list
		y := 688-((len(c.channels)*20)+29)
		g_game.GetGuiImage(IMG_GUI_CHATLISTTOP).Draw(0, y)
		
		//channels
		y = 688-(len(c.channels)*20)
		i := 0
		for _, channel := range c.channels {
			if  channel.closable {
				g_game.GetGuiImage(IMG_GUI_CHATLISTMIDX).Draw(0, y+(i*20))
			} else if channel.gamechannel {
				if channel.notifications {
					g_game.GetGuiImage(IMG_GUI_CHATLISTMIDLISTEN).Draw(0, y+(i*20))
				} else {
					g_game.GetGuiImage(IMG_GUI_CHATLISTMIDIGNORE).Draw(0, y+(i*20))
				}
			} else {
				g_game.GetGuiImage(IMG_GUI_CHATLISTMID).Draw(0, y+(i*20))
			}
			
			font := g_engine.GetFont(FONT_PURITANBOLD_14)
			if channel.updated && c.listBlink {
				hasUpdate = true
				font.SetColor(255, 0, 0)
			} else {
				font.SetColor(255, 255, 255)
			}
			
			font.DrawTextInRect(channel.name, 3, 3+y+(i*20), NewRect(3,y+(i*20),80,19))
			i++
		}
	}
	
	if c.listBlink && !hasUpdate {
		for _, channel := range c.channels {
			if channel.updated {
				hasUpdate = true
			}
		}
	}
	
	font := g_engine.GetFont(FONT_ARIALBLACK_10)
	if c.listBlink && hasUpdate {
		font.SetColor(255, 0 , 0)
	} else {
		font.SetColor(0, 0 ,0)
	}
	font.DrawText(c.channels[c.activeChannel].name, 28, 704)
}

func (c *PU_Chat) SendMessage(_text string) {
	if channel, exists := c.channels[c.activeChannel]; exists {
		if channel.id == CHANNEL_IRC {
			//send irc message
		} else if channel.gamechannel || channel.id == CHANNEL_LOCAL {
			g_conn.Game().SendChat(SPEAK_NORMAL, c.activeChannel, _text)
		} else {
			//private message 
			g_conn.Game().SendChat(SPEAK_PRIVATE, c.activeChannel, _text)
		}
	}
}

func (c *PU_Chat) Draw() {
	if !c.visible {
		return
	}
	c.DrawChatList()
}

func (c *PU_Chat) MouseDown(_x int, _y int) {

}

func (c *PU_Chat) MouseUp(_x int, _y int) {
	activeChannel := NewRect(20,694,66,19)
	if activeChannel.Contains(_x, _y) {
		c.listVisible = !c.listVisible
		return
	}
	
	if c.listVisible {
		h := ((len(c.channels)*20)+29)
		y := 688-h
		
		channelList := NewRect(0,y,116,h)
		if channelList.Contains(_x, _y) {
			y += 29
			i := 0
			for id, channel := range c.channels {
				channelName := NewRect(0,y+(i*20),82,20)
				channelClose := NewRect(90,3+y+(i*20),13,13)
				
				if channelName.Contains(_x, _y) {
					c.SetActive(id)
					c.listVisible = false
					return
				} else if channelClose.Contains(_x, _y) {
					if channel.closable {
						c.CloseChannel(id)
					} else if channel.gamechannel {
						if channel.notifications {
							channel.notifications = false
						} else {
							channel.notifications = true
						}
					}
					return
				}
				i++
			}
		}
	}
}

func (c *PU_Chat) MouseMove(_x int, _y int) {

}

func (c *PU_Chat) MouseScroll(_dir int) {

}

func (c *PU_Chat) Focusable() bool {
	return false
}

func (c *PU_Chat) KeyDown(_keysym int, _scancode int) {

}

