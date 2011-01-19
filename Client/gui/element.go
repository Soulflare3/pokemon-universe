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

type IGuiElement interface {
	Draw()
	
	MouseDown(_x int, _y int)
	MouseUp(_x int, _y int)
	MouseMove(_x int, _y int)
	MouseScroll(_dir int)
	KeyDown(_keysym int, _scancode int)
	
	GetRect() *PU_Rect
	IsVisible() bool
	SetParent(_parent IGuiElement)
	GetParent() IGuiElement
	SetFocus(_focus bool)
	HasFocus() bool
}

type PU_GuiElement struct {
	parent IGuiElement
	focus bool
	rect *PU_Rect
	visible bool
}

func (g *PU_GuiElement) SetRect(_rect *PU_Rect) {
	g.rect = _rect
}

func (g *PU_GuiElement) GetRect() *PU_Rect {
	return g.rect
}

func (g *PU_GuiElement) SetVisible(_visible bool) {
	g.visible = _visible
}

func (g *PU_GuiElement) IsVisible() bool {
	return g.visible
}

func (g *PU_GuiElement) SetParent(_parent IGuiElement) {
	g.parent = _parent
}

func (g *PU_GuiElement) GetParent() IGuiElement {
	return g.parent
}

func (g *PU_GuiElement) SetFocus(_focus bool) {
	g.focus = _focus
}

func (g *PU_GuiElement) HasFocus() bool {
	return g.focus
}
