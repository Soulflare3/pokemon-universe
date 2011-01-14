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
	list "container/list"
)

type IGuiElement interface {
	Draw()
	
	MouseDown(_x int, _y int)
	MouseUp(_x int, _y int)
	MouseMove(_x int, _y int)
	MouseScroll()
	KeyDown(_keysym int, _scancode int)
	
	GetRect() *PU_Rect
	IsVisible() bool
	SetParent(_element IGuiElement)
	GetParent() IGuiElement
}

type PU_Gui struct {
	elementList *list.List
}

func NewGui() *PU_Gui {
	return &PU_Gui{elementList : list.New()}
}

func (g *PU_Gui) AddElement(_element IGuiElement) {
	g.elementList.PushBack(_element)
}

func (g *PU_Gui) RemoveElement(_element IGuiElement) {
	for e := g.elementList.Front(); e != nil;  e = e.Next() {
		if e.Value == _element {
			g.elementList.Remove(e)
			break
		}
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
