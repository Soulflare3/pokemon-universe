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

type PU_ChatChannel struct {
	id int
	name string
	
	closable bool
	updated bool
	notifications bool
	
	gamechannel bool
	
	scrollbar *PU_Scrollbar
	textbox *PU_Textbox
}

const (
	CHANNEL_LOG = -2
	CHANNEL_IRC = -1
	CHANNEL_LOCAL = 0
	CHANNEL_WORLD = 1
	CHANNEL_TRADE = 2
	CHANNEL_BATTLE = 3
)

const (
	CHATCHANNEL_NAME_MAXWIDTH = 55
	
	CHAT_SCROLLBAR_X = 372
	CHAT_SCROLLBAR_Y = 594
	CHAT_SCROLLBAR_HEIGHT = 93
	
	CHAT_TEXTBOX_X = 0
	CHAT_TEXTBOX_Y = 591
	CHAT_TEXTBOX_WIDTH = 375
	CHAT_TEXTBOX_HEIGHT = 100
)

func NewChatChannel(_id int, _name string) *PU_ChatChannel {
	channel := &PU_ChatChannel{}
	channel.id = _id
	
	namefont := g_engine.GetFont(FONT_PURITANBOLD_14)
	namelen := 0
	for i := 0; i < len(_name); i++ {
		char := string(_name[i])
		charsize := namefont.GetStringWidth(char)
		if namelen+charsize <= 55 {
			channel.name += char
			namelen += charsize
		} else {
			channel.name += ".."
			break
		}
	}
	
	channel.closable = true	
	channel.notifications = true
	
	channel.scrollbar = NewScrollbar(CHAT_SCROLLBAR_X, CHAT_SCROLLBAR_Y, CHAT_SCROLLBAR_HEIGHT)
	channel.scrollbar.visible = false
	channel.scrollbar.maxvalue = 0
	
	channel.textbox = NewTextbox(NewRect(CHAT_TEXTBOX_X, CHAT_TEXTBOX_Y, CHAT_TEXTBOX_WIDTH, CHAT_TEXTBOX_HEIGHT), FONT_PURITANBOLD_12)
	channel.textbox.visible = false
	channel.textbox.scrollbar = channel.scrollbar
	
	switch _id {
		case CHANNEL_WORLD, CHANNEL_TRADE, CHANNEL_BATTLE, CHANNEL_IRC, CHANNEL_LOG:
			channel.gamechannel = true
	}
	
	return channel
}

func (c *PU_ChatChannel) AddMessage(_text *PU_Text) {
	if c.textbox != nil {
		c.textbox.AddText(_text)
	}
}

func (c *PU_ChatChannel) Close() {
	g_gui.RemoveElement(c.textbox)
	g_gui.RemoveElement(c.scrollbar)
}

func (c *PU_ChatChannel) SetActive(_active bool) {
	if c.textbox != nil && c.scrollbar != nil {
		if _active {
			c.textbox.visible = true
			c.scrollbar.visible = true
		} else {
			c.textbox.visible = false
			c.scrollbar.visible = false
		}
	}
}

