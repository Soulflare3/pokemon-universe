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
