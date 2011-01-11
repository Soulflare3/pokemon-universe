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
package position

import (
	"strconv"
)

type Position struct {
	X, Y, Z int
}

// ZP is the zero position
var ZP Position

// String returns a string represntation of p like "3,9,1"
func (p Position) String() string {
	return strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y) + "," + strconv.Itoa(p.Z)
}

// Add returns the position p+q.
func (p Position) Add(q Position) Position {
    return Position{p.X + q.X, p.Y + q.Y, p.Z + q.Z}
}

// Sub returns the position p-q.
func (p Position) Sub(q Position) Position {
    return Position{p.X - q.X, p.Y - q.Y, p.Z - q.Z}
}

// Eq returns true if Position p and q are the same
func (p Position) Eq(_q Position) bool {
	return (p.X == _q.X) && (p.Y == _q.Y) && (p.Z == _q.Z)
}

// Equals calls and returns outcome of Eq
func (p Position) Equals(_q Position) bool {
	return p.Eq(_q)
}
