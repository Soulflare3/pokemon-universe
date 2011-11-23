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
	"math"
)

const (
	ONSCREENCHAT_TICKS = 3000
)


type PU_OnscreenChatLine struct {
	text  string
	ticks int
}

var lolr, lolg, lolb int = 0, 162, 232

func NewOnscreenChatLine(_text string) *PU_OnscreenChatLine {
	return &PU_OnscreenChatLine{text: _text, ticks: ONSCREENCHAT_TICKS}
}

type PU_OnscreenChatMessage struct {
	name  string
	x     int
	y     int
	lines []*PU_OnscreenChatLine
}

func NewOnscreenChatMessage(_name string, _x int, _y int) *PU_OnscreenChatMessage {
	return &PU_OnscreenChatMessage{name: _name, x: _x, y: _y}
}

func NewOnscreenChatMessageExt(_name string, _x int, _y int, _text string) *PU_OnscreenChatMessage {
	msg := NewOnscreenChatMessage(_name, _x, _y)
	msg.AddText(_text)
	return msg
}

func (m *PU_OnscreenChatMessage) Draw(_ticks int) bool {
	ret := true

	offsetX, offsetY := g_game.GetScreenOffset()

	m.UpdateLines(_ticks)

	font := g_engine.GetFont(FONT_PURITANBOLD_14)
	if len(m.lines) > 0 {
		lineHeight := font.GetStringHeight()
		height := lineHeight + (lineHeight * len(m.lines))

		center := false

		drawX := MID_X - (int(g_game.self.GetX()) - m.x)
		drawY := MID_Y - (int(g_game.self.GetY()) - m.y)

		drawX = (drawX * TILE_WIDTH) - TILE_WIDTH - 22 + offsetX
		drawY = (drawY * TILE_HEIGHT) - TILE_HEIGHT + offsetY

		switch {
		case drawY-height < 0:
			drawY = 0

		case drawY > WINDOW_HEIGHT:
			drawY = WINDOW_HEIGHT - height

		default:
			drawY -= height
			drawY += lineHeight
		}

		header := m.name + " says:"

		widest := font.GetStringWidth(header)

		for i := 0; i < len(m.lines); i++ {
			line := m.lines[i]
			if line != nil {
				len := font.GetStringWidth(line.text)
				if len > widest {
					widest = len
				}
			}
		}

		switch {
		case (drawX - int(math.Ceil(float64(widest)/2.0))) < 0:
			drawX = 0

		case (drawX + int(math.Ceil(float64(widest)/2.0))) > WINDOW_WIDTH:
			drawX = WINDOW_WIDTH - widest

		default:
			center = true
		}

		posHalf := 0
		if !center {
			posHalf = drawX + int(math.Ceil(float64(((drawX+widest)-drawX)/2)))
		} else {
			posHalf = (drawX - (int(math.Ceil(float64(widest) / 2.0)))) + int(math.Ceil(float64(((drawX+widest)-(drawX-(int(math.Ceil(float64(widest)/2.0)))))/2)))
		}

		nameHalf := int(math.Floor(float64(font.GetStringWidth(header) / 2.0)))
		centerPos := posHalf - nameHalf

		font.SetColor(187, 99, 245)

		font.SetStyle(true, false, false)
		font.DrawBorderedText(header, centerPos, drawY)

		for i := 0; i < len(m.lines); i++ {
			line := m.lines[i]
			if line != nil {
				nameHalf = int(math.Floor(float64(font.GetStringWidth(line.text)) / 2.0))
				centerPos = posHalf - nameHalf

				//font.DrawBorderedText(line.text, centerPos, (drawY+height)-((i+1)*lineHeight))
				font.DrawBorderedText(line.text, centerPos, drawY+((i+1)*lineHeight))
			}
		}
	} else {
		ret = false
	}
	return ret
}

func (m *PU_OnscreenChatMessage) AddLine(_text string) {
	m.lines = append(m.lines, NewOnscreenChatLine(_text))
	if len(m.lines) > 4 {
		m.lines = append(m.lines[:0], m.lines[1:]...)
	}
}

func (m *PU_OnscreenChatMessage) AddText(_text string) {
	font := g_engine.GetFont(FONT_PURITANBOLD_14)
	curSize := 2
	curText := ""
	maxWidth := 160
	textWidth := font.GetStringWidth(_text) + 2

	if textWidth > maxWidth {
		text := _text
		curPos := 0
		for curPos < len(text) {
			word := m.NextWord(text, curPos)
			wordSize := font.GetStringWidth(word)
			if curSize+wordSize < maxWidth {
				curText += word
				curSize += wordSize
				curPos += len(word)
			} else {
				if curText != "" {
					m.AddLine(curText)

					curText = ""
					curSize = 2
				} else {
					for i := 0; i < len(word); i++ {
						charWidth := font.GetStringWidth(string(word[i]))
						if curSize+charWidth > maxWidth {
							curText += "-"

							m.AddLine(curText)

							curText = ""
							curSize = 2

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
			m.AddLine(curText)
		}
	} else {
		m.AddLine(_text)
	}
}

func (t *PU_OnscreenChatMessage) NextWord(_text string, _start int) string {
	for i := _start; i < len(_text); i++ {
		if _text[i] == ' ' {
			return string(_text[_start : i+1])
		}
	}
	return string(_text[_start:])
}

func (m *PU_OnscreenChatMessage) UpdateLines(_ticks int) {
	for i := 0; i < len(m.lines); {
		line := m.lines[i]
		if line != nil {
			line.ticks -= _ticks
			if line.ticks <= 0 {
				m.lines = append(m.lines[:i], m.lines[i+1:]...)
			} else {
				i++
			}
		}
	}
}
