package main

import (
	"fmt"
	"sdl"
)

type PU_Rect struct {
	x, y, width, height int
}

func NewRect(_x int, _y int, _width int, _height int) *PU_Rect {
	return &PU_Rect{x : _x,
					y : _y,
					width : _width,
					height : _height}
}

func (r *PU_Rect) Equals(_rect *PU_Rect) bool {
	return ((r.x == _rect.x) && (r.y == _rect.y) && 
			(r.width == _rect.width) && (r.height == _rect.height))
}

func (r *PU_Rect) Contains(_x int, _y int) bool {
	return (_x >= r.x && _x <= r.x+r.width && _y >= r.y && _y <= r.y+r.height)
}

func (r *PU_Rect) ContainsRect(_rect *PU_Rect) bool {
	return (r.Contains(_rect.x,_rect.y) && 
			r.Contains(_rect.x+_rect.width,_rect.y) && 
			r.Contains(_rect.x,_rect.y+_rect.height) && 
			r.Contains(_rect.x+_rect.width,_rect.y+_rect.height))
}

func (r *PU_Rect) Intersects(_rect *PU_Rect) bool {
	return !(r.x > _rect.x+_rect.width || _rect.x > r.x+r.width ||
			 r.y > _rect.y+_rect.height || _rect.y > r.y+r.height)
}

func (r *PU_Rect) Intersection(_rect *PU_Rect) *PU_Rect {
	if r.Intersects(_rect) {
		tempX := _rect.x
		if r.x > _rect.x {
			tempX = r.x
		}
		
		tempY := _rect.y
		if r.y > _rect.y {
			tempY = r.y
		}
		
		tempW := _rect.x+_rect.width
		if r.x+r.width < _rect.x+_rect.width {
			tempW = r.x+r.width
		}
		
		tempH := _rect.y+_rect.height
		if r.y+r.height < _rect.y+_rect.height {
			tempH = r.y+r.height
		}
		
		tempW -= tempX
		tempH -= tempY
		
		return &PU_Rect{x : tempX,
						y : tempY,
						width : tempW,
						height : tempH}
	}
	return nil
}

func (r *PU_Rect) ToSDL() *sdl.Rect {
	return &sdl.Rect{int32(r.x), int32(r.y), int32(r.width), int32(r.height)}
}

func (r *PU_Rect) ToString() string {
	return fmt.Sprintf("PU_Rect(%v,%v,%v,%v)",r.x, r.y, r.width, r.height)
}

