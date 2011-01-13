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
	GetID() int
	Draw()
	
	MouseDown(_x int, _y int)
	MouseUp(_x int, _y int)
	MouseMove(_x int, _y int)
	MouseScroll()
	KeyDown(_keysym int, _scancode int)
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
