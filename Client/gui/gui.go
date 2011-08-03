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

type PU_Gui struct {
	elementList []IGuiElement
}

func NewGui() *PU_Gui {
	return &PU_Gui{elementList: make([]IGuiElement, 0)}
}

func (g *PU_Gui) AddElement(_element IGuiElement) {
	g.elementList = append(g.elementList, _element)
}

func (g *PU_Gui) ElementIndex(_element IGuiElement) int {
	for index, element := range g.elementList {
		if element == _element {
			return index
		}
	}
	return 0
}

func (g *PU_Gui) RemoveElement(_element IGuiElement) {
	a := make([]IGuiElement, len(g.elementList)-1)
	i := 0
	for _, element := range g.elementList {
		if element != _element {
			a[i] = element
			i++
		}
	}
	g.elementList = a
}

//Helper method to draw GUI images (at the right location and possibly clipped)
func (g *PU_Gui) DrawImage(_element IGuiElement, _image *PU_Image, _src *PU_Rect, _dst *PU_Rect) {
	if parent := _element.GetParent(); parent != nil {
		parentRect := NewRect(0, 0, parent.GetRect().width, parent.GetRect().height)
		if parentRect.ContainsRect(_dst) || parentRect.Intersects(_dst) {
			clipRect := parentRect.Intersection(_dst)
			clipRect.x += parent.GetRect().x
			clipRect.y += parent.GetRect().y

			g.DrawImage(parent, _image, _src, clipRect)
		}
	} else {
		if _dst.width == 0 && _dst.height == 0 {
			_dst.width = WINDOW_WIDTH
			_dst.height = WINDOW_HEIGHT
		}

		_image.DrawRectInRect(_src, _dst)
	}
}

//Find out within what frame (represented by a rect) the GUI element will be drawn
func (g *PU_Gui) GetClipRect(_element IGuiElement, _rect *PU_Rect) *PU_Rect {
	if parent := _element.GetParent(); parent != nil {
		parentRect := NewRect(0, 0, parent.GetRect().width, parent.GetRect().height)
		if (parentRect.ContainsRect(_rect)) || (parentRect.Intersects(_rect)) {
			clipRect := parentRect.Intersection(_rect)
			g.GetClipRect(parent, clipRect)
		}
	} else {
		if _rect.width == 0 && _rect.height == 0 {
			_rect.width = WINDOW_WIDTH
			_rect.height = WINDOW_HEIGHT
		}
		return _rect
	}
	return nil //keep the compiler happy
}

//Get the coordinates of the root element of an element 
func (g *PU_Gui) GetTopRect(_element IGuiElement, _rect *PU_Rect) *PU_Rect {
	if parent := _element.GetParent(); parent != nil {
		_rect.x += parent.GetRect().x
		_rect.y += parent.GetRect().y
		g.GetTopRect(parent, _rect)
	} else {
		return _rect
	}
	return nil
}

func (g *PU_Gui) SetFocus(_element IGuiElement) {
	//remove all focuses and set it for the arg element
	for _, e := range g.elementList {
		if e == _element {
			e.SetFocus(true)
		} else {
			e.SetFocus(false)
		}
	}
}

func (g *PU_Gui) NextFocus() {
	foundPos := false
	for i := 0; i < 2; i++ {
		for _, element := range g.elementList {
			if !foundPos && element.HasFocus() {
				element.SetFocus(false)
				foundPos = true
			} else if foundPos && element.Focusable() {
				element.SetFocus(true)
				return
			}
		}
	}
}

func (g *PU_Gui) Draw() {
	for _, e := range g.elementList {
		e.Draw()
	}
}

func (g *PU_Gui) MouseDown(_x int, _y int) {
	for _, e := range g.elementList {
		e.MouseDown(_x, _y)
	}
}

func (g *PU_Gui) MouseUp(_x int, _y int) {
	for _, e := range g.elementList {
		e.MouseUp(_x, _y)
	}
}

func (g *PU_Gui) MouseMove(_x int, _y int) {
	for _, e := range g.elementList {
		e.MouseMove(_x, _y)
	}
}

func (g *PU_Gui) MouseScroll(_dir int) {
	for _, e := range g.elementList {
		e.MouseScroll(_dir)
	}
}

func (g *PU_Gui) KeyDown(_keysym int, _scancode int) {
	for _, e := range g.elementList {
		if _scancode == 9 && e.HasFocus() {
			e.KeyDown(_keysym, _scancode)
			return
		}
		e.KeyDown(_keysym, _scancode)
	}
}
