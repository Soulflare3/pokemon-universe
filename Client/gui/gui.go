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
	for e := g.elementList.Front(); e != nil; e = e.Next() {
		if e.Value == _element {
			e.Value.(IGuiElement).SetFocus(true)
		} else {
			e.Value.(IGuiElement).SetFocus(false)
		}
	}
}

func (g *PU_Gui) NextFocus() {
	//find the current focus holder and remove all focuses
	var currentFocus *list.Element = nil
	for e := g.elementList.Front(); e != nil;  e = e.Next() {
		if e.Value.(IGuiElement).HasFocus() {
			e.Value.(IGuiElement).SetFocus(false)
			currentFocus = e
		}
	}
	
	//if there's no current focus holder we can just start at the front
	if currentFocus == nil {
		currentFocus = g.elementList.Front()
	} else {
		currentFocus = currentFocus.Next()
	}
	
	//find a new focusable element (starting at the current element's list position)
	for e := currentFocus; e != nil; e = e.Next() {
		if e.Value.(IGuiElement).Focusable() {
			e.Value.(IGuiElement).SetFocus(true)
			return
		}
	}
	
	//reaching this means we haven't found any elements to the left of the current element
	
	//if we already searched the whole list, we just give up here
	if currentFocus == g.elementList.Front() {
		return
	}
	
	//search the list again, starting at the front
	for e := g.elementList.Front(); e != nil; e = e.Next() {
		if e.Value.(IGuiElement).Focusable() {
			e.Value.(IGuiElement).SetFocus(true)
			return
		}
	}
}

func (g *PU_Gui) Draw() {
	for e := g.elementList.Front(); e != nil;  e = e.Next() {
		e.Value.(IGuiElement).Draw()
	}
}

func (g *PU_Gui) MouseDown(_x int, _y int) {
	for e := g.elementList.Front(); e != nil;  e = e.Next() {
		e.Value.(IGuiElement).MouseDown(_x, _y)
	}
}

func (g *PU_Gui) MouseUp(_x int, _y int) {
	for e := g.elementList.Front(); e != nil;  e = e.Next() {
		e.Value.(IGuiElement).MouseUp(_x, _y)
	}
}

func (g *PU_Gui) MouseMove(_x int, _y int) {
	for e := g.elementList.Front(); e != nil;  e = e.Next() {
		e.Value.(IGuiElement).MouseMove(_x, _y)
	}
}

func (g *PU_Gui) MouseScroll(_dir int) {
	for e := g.elementList.Front(); e != nil;  e = e.Next() {
		e.Value.(IGuiElement).MouseScroll(_dir)
	}
}

func (g *PU_Gui) KeyDown(_keysym int, _scancode int) {
	for e := g.elementList.Front(); e != nil;  e = e.Next() {
		if _scancode == 9 && e.Value.(IGuiElement).HasFocus() { //1 tab per event to avoid tab loop
			e.Value.(IGuiElement).KeyDown(_keysym, _scancode)
			return
		}
		e.Value.(IGuiElement).KeyDown(_keysym, _scancode)
	}
}

