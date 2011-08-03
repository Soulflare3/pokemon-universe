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

//PU_Textbox is a multiline textfield control that uses the PU_Text type as text. It can not be edited by the user.

import (
	"sdl"
	list "container/vector"
)

const (
	TEXTBOX_BUFFERSIZE = 100
)

type PU_Textbox struct {
	PU_GuiElement
	transparent bool
	font        *PU_Font
	bgcolor     sdl.Color

	bold       bool
	italic     bool
	underlined bool

	lines     list.Vector
	scrollbar *PU_Scrollbar
}

func NewTextbox(_rect *PU_Rect, _font int) *PU_Textbox {
	textbox := &PU_Textbox{transparent: true,
		font: g_engine.GetFont(_font)}
	textbox.rect = _rect
	textbox.visible = true
	g_gui.AddElement(textbox)
	return textbox
}

func (t *PU_Textbox) SetFont(_id int) {
	t.font = g_engine.GetFont(_id)
}

func (t *PU_Textbox) SetBgColor(_red uint8, _green uint8, _blue uint8) {
	t.bgcolor.R = _red
	t.bgcolor.G = _green
	t.bgcolor.B = _blue
}

func (t *PU_Textbox) SetStyle(_bold bool, _italic bool, _underlined bool) {
	t.bold, t.italic, t.underlined = _bold, _italic, _underlined
}

func (t *PU_Textbox) AddLine(_line *PU_Text) {
	if t.lines.Len() > TEXTBOX_BUFFERSIZE {
		t.lines.Delete(0)
	}
	t.lines.Push(_line)

	t.UpdateScrollbar()
}

func (t *PU_Textbox) AddText(_text *PU_Text) {
	curSize := 0
	curText := ""
	newText := NewTextWithFont(_text.font)
	maxWidth := t.rect.width - 6

	if _text.GetWidth() > maxWidth {
		for part := 0; part < _text.count; part++ {
			text := _text.GetPart(part).text
			curPos := 0
			for curPos < len(text) {
				word := t.NextWord(text, curPos)
				wordSize := t.font.GetStringWidth(word)
				if curSize+wordSize < maxWidth {
					curText += word
					curSize += wordSize
					curPos += len(word)
				} else {
					if curText != "" {
						newText.Add(curText, _text.GetPart(part).color)
						t.AddLine(newText)
						newText = NewTextWithFont(_text.font)

						curText = ""
						curSize = 0
					} else {
						for i := 0; i < len(word); i++ {
							charWidth := t.font.GetStringWidth(string(word[i]))
							if curSize+charWidth > maxWidth {
								curText += "-"

								newText.Add(curText, _text.GetPart(part).color)
								t.AddLine(newText)
								newText = NewTextWithFont(_text.font)

								curText = ""
								curSize = 0

								curPos += i

								break
							}
							curText += string(word[i])
							curSize += charWidth
						}
					}
				}
			}
			if curText != "" {
				newText.Add(curText, _text.GetPart(part).color)
				curText = ""
				if part+1 >= _text.count {
					t.AddLine(newText)
				}
			}
		}
	} else {
		t.AddLine(_text)
	}
}

func (t *PU_Textbox) NextWord(_text string, _start int) string {
	for i := _start; i < len(_text); i++ {
		if _text[i] == ' ' {
			return string(_text[_start : i+1])
		}
	}
	return string(_text[_start:])
}

func (t *PU_Textbox) UpdateScrollbar() {
	if t.scrollbar != nil {
		fontHeight := t.font.GetStringHeight()
		boxHeight := t.rect.height - 6
		visibleLines := int(float32(boxHeight) / float32(fontHeight))

		max := t.lines.Len() - visibleLines
		if max <= 0 {
			max = 0
		}

		if t.scrollbar.value == t.scrollbar.maxvalue {
			t.scrollbar.maxvalue = max
			t.scrollbar.value = max
		} else {
			t.scrollbar.maxvalue = max
		}
	}
}

func (t *PU_Textbox) Draw() {
	if !t.visible {
		return
	}

	if !t.transparent {
		g_engine.DrawFillRect(t.rect, &t.bgcolor, 200)
	}

	fontHeight := t.font.GetStringHeight()
	boxHeight := t.rect.height - 6
	visibleLines := int(float32(boxHeight) / float32(fontHeight))
	scrollInc := 0
	if t.scrollbar != nil {
		scrollInc = t.scrollbar.value
	}

	clip := NewRectFrom(t.rect)
	g_gui.GetClipRect(t, clip)

	top := NewRectFrom(t.rect)
	g_gui.GetTopRect(t, top)

	drawX := top.x + 3
	drawY := top.y + 3

	for line := scrollInc; line < (visibleLines + scrollInc); line++ {
		if line < t.lines.Len() {
			text, ok := t.lines.At(line).(*PU_Text)
			if ok {
				numParts := text.count
				for part := 0; part < numParts; part++ {
					curPart := text.GetPart(part)

					t.font.SetStyle(t.bold, t.italic, t.underlined)
					color := ColorKeyToSDL(curPart.color)
					t.font.SetColor(color.R, color.G, color.B)
					t.font.DrawTextInRect(curPart.text, drawX, drawY, clip)

					drawX += t.font.GetStringWidth(curPart.text)
				}
				drawX = top.x + 3
				drawY += fontHeight
			}
		}
	}
}

func (t *PU_Textbox) MouseDown(_x int, _y int) {

}

func (t *PU_Textbox) MouseUp(_x int, _y int) {

}

func (t *PU_Textbox) MouseMove(_x int, _y int) {

}

func (t *PU_Textbox) MouseScroll(_dir int) {
	if t.scrollbar == nil {
		return
	}

	curValue := t.scrollbar.value

	if _dir == sdl.SCROLL_UP {
		if curValue-1 >= 0 {
			t.scrollbar.value = curValue - 1
		}
	} else if _dir == sdl.SCROLL_DOWN {
		maxValue := t.scrollbar.maxvalue
		if curValue+1 <= maxValue {
			t.scrollbar.value = curValue + 1
		}
	}
}

func (t *PU_Textbox) Focusable() bool {
	return false
}

func (t *PU_Textbox) KeyDown(_keysym int, _scancode int) {

}
