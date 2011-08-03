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

const (
	SCROLLBTN_HEIGHT = 7
	SCROLLBTN_WIDTH  = 9
	SCROLLER_HEIGHT  = 8
)

type PU_Scrollbar struct {
	PU_GuiElement

	value     int
	minvalue  int
	maxvalue  int
	scrolling bool

	ValueChangedCallback func(_item int)
}

func NewScrollbar(_x int, _y int, _length int) *PU_Scrollbar {
	scrollbar := &PU_Scrollbar{}
	scrollbar.rect = NewRect(_x, _y, SCROLLBTN_WIDTH, _length)
	scrollbar.visible = true
	scrollbar.value = 0
	scrollbar.minvalue = 0
	scrollbar.maxvalue = 100
	g_gui.AddElement(scrollbar)
	return scrollbar
}

func (s *PU_Scrollbar) GetScrollerPos() (x int, y int) {
	barlength := s.rect.height - SCROLLBTN_HEIGHT - SCROLLBTN_HEIGHT - (SCROLLER_HEIGHT / 2)
	perc := float32(s.value) / float32(s.maxvalue)
	pos := int(perc * float32(barlength))

	x = s.rect.x
	y = s.rect.y + SCROLLBTN_HEIGHT + pos
	return
}

func (s *PU_Scrollbar) CheckValueByPos(_x int, _y int) {
	if _x >= s.rect.x && _x <= s.rect.x+s.rect.width {
		if _y > s.rect.y+SCROLLBTN_HEIGHT+2 && _y < s.rect.y+s.rect.height-SCROLLBTN_HEIGHT-2 { //scroll area
			barlength := s.rect.height - SCROLLBTN_HEIGHT - SCROLLBTN_HEIGHT - SCROLLER_HEIGHT
			pos := _y - (s.rect.y + SCROLLBTN_HEIGHT + (SCROLLER_HEIGHT / 2))
			perc := float32(pos) / float32(barlength)
			s.value = int(perc * float32(s.maxvalue))
		} else if _y > s.rect.y+SCROLLBTN_HEIGHT && _y <= s.rect.y+SCROLLBTN_HEIGHT+2 {
			s.value = 0
		} else if _y >= s.rect.y+s.rect.height-SCROLLBTN_HEIGHT-2 && _y < s.rect.y+s.rect.height-SCROLLBTN_HEIGHT {
			s.value = s.maxvalue
		}
	}
}

func (s *PU_Scrollbar) Draw() {
	if !s.visible {
		return
	}

	//find the correct coordinates to draw at (when the element is embedded in another)
	top := NewRectFrom(s.rect)
	g_gui.GetTopRect(s, top)

	//top button
	g_gui.DrawImage(s, g_game.GetGuiImage(IMG_GUI_VSCROLLBAR_UP), NewRect(top.x, top.y, SCROLLBTN_WIDTH, SCROLLBTN_HEIGHT), s.rect)

	//mid body
	g_gui.DrawImage(s, g_game.GetGuiImage(IMG_GUI_VSCROLLBAR_MID), NewRect(top.x, top.y+SCROLLBTN_HEIGHT, SCROLLBTN_WIDTH, s.rect.height-(2*SCROLLBTN_HEIGHT)), s.rect)

	//bottom button
	g_gui.DrawImage(s, g_game.GetGuiImage(IMG_GUI_VSCROLLBAR_DOWN), NewRect(top.x, top.y+(s.rect.height-SCROLLBTN_HEIGHT), SCROLLBTN_WIDTH, SCROLLBTN_HEIGHT), s.rect)

	//scroller
	if s.maxvalue > 0 {
		x, y := s.GetScrollerPos()

		g_gui.DrawImage(s, g_game.GetGuiImage(IMG_GUI_VSCROLLBAR_SCROLLER), NewRect(x+1, y, SCROLLBTN_HEIGHT, SCROLLER_HEIGHT), s.rect)
	}
}

func (s *PU_Scrollbar) MouseDown(_x int, _y int) {
	if _x >= s.rect.x && _x <= s.rect.x+s.rect.width {
		_, sy := s.GetScrollerPos()

		if _y >= sy && _y <= sy+SCROLLER_HEIGHT {
			s.scrolling = true
		}
	}
}

func (s *PU_Scrollbar) MouseUp(_x int, _y int) {
	if s.scrolling {
		s.scrolling = false
	} else {
		s.CheckValueByPos(_x, _y)
	}

	if _x >= s.rect.x && _x <= s.rect.x+s.rect.width {
		if _y >= s.rect.y && _y <= s.rect.y+SCROLLBTN_HEIGHT {
			if s.value > 0 {
				s.value--
			}
		} else if _y >= s.rect.y+s.rect.height-SCROLLBTN_HEIGHT && _y <= s.rect.y+s.rect.height {
			if s.value < s.maxvalue {
				s.value++
			}
		}
	}
}

func (s *PU_Scrollbar) MouseMove(_x int, _y int) {
	if s.scrolling {
		s.CheckValueByPos(_x, _y)
	}
}

func (s *PU_Scrollbar) MouseScroll(_dir int) {

}

func (s *PU_Scrollbar) Focusable() bool {
	return false
}

func (s *PU_Scrollbar) KeyDown(_keysym int, _scancode int) {

}
