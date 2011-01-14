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

import "sdl"

type PU_Textfield struct {
	rect *PU_Rect
	visible bool
	transparent bool //when true, only the inserted text and the caret is visible
	font *PU_Font
	color sdl.Color
	text string
	password bool
}

func NewTextfield(_rect *PU_Rect, _font int) *PU_Textfield {
	textfield := &PU_Textfield{rect : _rect,
							   visible : true,
							   transparent : true,
							   font : g_engine.GetFont(_font),
							   password : false}
	textfield.SetColor(0, 0, 0)
	g_gui.AddElement(textfield)
	return textfield
}

func (t *PU_Textfield) SetFont(_id int) {
	t.font = g_engine.GetFont(_id)
}

func (t *PU_Textfield) SetColor(_red uint8, _green uint8, _blue uint8)  {
	t.color.R = _red
	t.color.G = _green
	t.color.B = _blue
}

func (t *PU_Textfield) Draw() {
	if !t.visible {
		return
	}
	
	if !t.transparent {
		//draw element ui here
	}
	
	//draw text
	if len(t.text) > 0 {
		var drawText string
		if t.password {
			for i := 0; i < len(t.text); i++ {
				drawText += "*"
			}
		} else {
			drawText = t.text
		}
		
		//find the correct coordinates to draw at (when the element is embedded in another)
		top := NewRect(t.rect.x, t.rect.y, t.rect.width, t.rect.height)
		g_gui.GetTopRect(t, top)
		
		//find out within what frame the text should be drawn 		
		clip := NewRect(t.rect.x, t.rect.y, t.rect.width, t.rect.height)
		g_gui.GetClipRect(t, clip)
		
		t.font.SetColor(t.color.R, t.color.G, t.color.B)
		t.font.DrawTextInRect(drawText, top.x, top.y, clip)
	}
}

func (t *PU_Textfield)MouseDown(_x int, _y int) {

}

func (t *PU_Textfield) MouseUp(_x int, _y int) {

}

func (t *PU_Textfield)MouseMove(_x int, _y int) {

}

func (t *PU_Textfield)MouseScroll() {

}

func (t *PU_Textfield)KeyDown(_keysym int, _scancode int) {

}
	
func (t *PU_Textfield) GetRect() *PU_Rect {
	return nil
}

func (t *PU_Textfield)IsVisible() bool {
	return t.visible
}

func (t *PU_Textfield)SetParent(_element IGuiElement) {

}

func (t *PU_Textfield)GetParent() IGuiElement {
	return nil
}

