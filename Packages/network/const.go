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

package network

const (
	HEADER_PING				= 0x00
	HEADER_LOGIN  			= 0x01
	HEADER_LOGOUT 			= 0x02
	
	HEADER_IDENTITY			= 0xAA
	
	HEADER_WALK				= 0xB1
	HEADER_CANCELWALK		= 0xB2
	HEADER_WARP				= 0xB3
	HEADER_TURN				= 0xB4
	
	HEADER_TILES			= 0xC1
	HEADER_ADDCREATURE		= 0xC2
	HEADER_REMOVECREATURE	= 0xC3
	
	
	// These need to get other values
	HEADER_REFRESHCOMPLETE	= 0x03
	HEADER_REFRESHWORLD		= 0xC4
)
