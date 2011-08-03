/*Pokemon Universe MMORPG
Copyright (C) 2010 the Pokemon Universe Authors

This program is free software you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.*/
package main

func ClipText(_text string, _font int, _maxwidth int) []string {
	font := g_engine.GetFont(_font)
	if font == nil {
		return nil
	}

	curSize := 2
	curText := ""
	maxWidth := _maxwidth
	textWidth := font.GetStringWidth(_text) + 2
	ret := make([]string, 0)

	if textWidth > maxWidth {
		text := _text
		curPos := 0
		for curPos < len(text) {
			word := NextWord(text, curPos)
			wordSize := font.GetStringWidth(word)
			if curSize+wordSize < maxWidth {
				curText += word
				curSize += wordSize
				curPos += len(word)
			} else {
				if curText != "" {
					ret = append(ret, curText)

					curText = ""
					curSize = 2
				} else {
					for i := 0; i < len(word); i++ {
						charWidth := font.GetStringWidth(string(word[i]))
						if curSize+charWidth > maxWidth {
							curText += "-"

							ret = append(ret, curText)

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
			ret = append(ret, curText)
		}
	} else {
		ret = append(ret, _text)
	}
	return ret
}

func NextWord(_text string, _start int) string {
	for i := _start; i < len(_text); i++ {
		if _text[i] == ' ' {
			return string(_text[_start : i+1])
		}
	}
	return string(_text[_start:])
}

func DrawType(_type string, _x int, _y int) {
	var img *PU_Image
	switch _type {
	case "ground":
		img = nil
	case "water":
		img = g_game.GetGuiImage(IMG_GUI_TYPEWATER)
	case "ghost":
		img = g_game.GetGuiImage(IMG_GUI_TYPEGHOST)
	case "bug":
		img = nil
	case "fighting":
		img = g_game.GetGuiImage(IMG_GUI_TYPEFIGHTING)
	case "psychic":
		img = g_game.GetGuiImage(IMG_GUI_TYPEPSYCHIC)
	case "grass":
		img = g_game.GetGuiImage(IMG_GUI_TYPEGRASS)
	case "dark":
		img = nil
	case "normal":
		img = g_game.GetGuiImage(IMG_GUI_TYPENORMAL)
	case "poison":
		img = nil
	case "electric":
		img = nil
	case "unknown":
		img = nil
	case "steel":
		img = nil
	case "rock":
		img = g_game.GetGuiImage(IMG_GUI_TYPEROCK)
	case "dragon":
		img = nil
	case "flying":
		img = g_game.GetGuiImage(IMG_GUI_TYPEFLYING)
	case "fire":
		img = g_game.GetGuiImage(IMG_GUI_TYPEFIRE)
	case "ice":
		img = nil
	default:
		img = nil //unknown
	}

	if img == nil {
		img = g_game.GetGuiImage(IMG_GUI_TYPENORMAL)
	}

	img.Draw(_x, _y)
}
