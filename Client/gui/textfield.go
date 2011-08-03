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
	"fmt"
)

type PU_Textfield struct {
	PU_GuiElement
	transparent bool //when true, only the inserted text and the caret is visible
	font        *PU_Font
	color       sdl.Color
	bgcolor     sdl.Color
	text        string
	password    bool
	caret       bool
	caretLast   uint32
	readonly    bool

	bold       bool
	italic     bool
	underlined bool

	KeyDownCallback func(_keysym int, _scancode int)
}

func NewTextfield(_rect *PU_Rect, _font int) *PU_Textfield {
	textfield := &PU_Textfield{transparent: true,
		font:      g_engine.GetFont(_font),
		caretLast: sdl.GetTicks()}
	textfield.rect = _rect
	textfield.visible = true
	textfield.SetColor(0, 0, 0)
	g_gui.AddElement(textfield)
	return textfield
}

func (t *PU_Textfield) SetFont(_id int) {
	t.font = g_engine.GetFont(_id)
}

func (t *PU_Textfield) SetColor(_red uint8, _green uint8, _blue uint8) {
	t.color.R = _red
	t.color.G = _green
	t.color.B = _blue
}

func (t *PU_Textfield) SetBgColor(_red uint8, _green uint8, _blue uint8) {
	t.bgcolor.R = _red
	t.bgcolor.G = _green
	t.bgcolor.B = _blue
}

func (t *PU_Textfield) SetStyle(_bold bool, _italic bool, _underlined bool) {
	t.bold, t.italic, t.underlined = _bold, _italic, _underlined
}

func (t *PU_Textfield) Draw() {
	if !t.visible {
		return
	}

	if !t.transparent {
		g_engine.DrawFillRect(t.rect, &t.bgcolor, 200)
	}

	var caretX int

	//find the correct coordinates to draw at (when the element is embedded in another)
	top := NewRectFrom(t.rect)
	g_gui.GetTopRect(t, top)

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

		//find out within what frame the text should be drawn 		
		clip := NewRectFrom(t.rect)
		g_gui.GetClipRect(t, clip)

		t.font.SetStyle(t.bold, t.italic, t.underlined)
		t.font.SetColor(t.color.R, t.color.G, t.color.B)
		t.font.DrawTextInRect(drawText, top.x, top.y, clip)

		caretX = top.x + t.font.GetStringWidth(drawText)
	}

	//draw caret
	if t.focus && t.caret && !t.readonly {
		image := g_game.GetGuiImage(IMG_GUI_CARET)
		if image != nil {
			if caretX == 0 {
				caretX = top.x
			}

			lineheight := t.font.GetStringHeight()
			drawRect := NewRect(caretX, top.y+1, int(image.w), lineheight)
			g_gui.DrawImage(t, image, drawRect, t.rect)
		}
	}
	if sdl.GetTicks()-t.caretLast >= 500 {
		t.caret = !t.caret
		t.caretLast = sdl.GetTicks()
	}
}

func (t *PU_Textfield) MouseDown(_x int, _y int) {

}

func (t *PU_Textfield) MouseUp(_x int, _y int) {
	if t.rect.Contains(_x, _y) && !t.readonly {
		g_gui.SetFocus(t)
	}
}

func (t *PU_Textfield) MouseMove(_x int, _y int) {

}

func (t *PU_Textfield) MouseScroll(_dir int) {

}

func (t *PU_Textfield) Focusable() bool {
	if !t.readonly {
		return true
	}
	return false
}

func (t *PU_Textfield) KeyDown(_keysym int, _scancode int) {
	if !t.focus || t.readonly {
		return
	}

	if _scancode == 8 { //backspace
		if len(t.text) > 0 {
			textlength := len(t.text)
			t.text = t.text[0 : textlength-1]
		}
	} else if _scancode == 9 { //tab
		g_gui.NextFocus()
	} else if _keysym != 0 && _keysym > 31 { //normal text input
		t.text += fmt.Sprintf("%c", _keysym)
	}

	if t.KeyDownCallback != nil {
		t.KeyDownCallback(_keysym, _scancode)
	}
}
