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
package db

import (
	"os"
)

type ClientInfo struct {
	SQLHost string
	SQLUser string
	SQLPass string
	SQLDB   string
}

type Database interface {
	// Connect to a database
	Connect(_client ClientInfo) (err os.Error)

	// Close connection
	Close()

	// Execute a query without storing the result
	// Will return an error when failed
	ExecuteQuery(_query string) (err os.Error)

	// Execute a query and store the output in a resultset
	// Will return an error when failed
	StoreQuery(_query string) (res ResultSet, err os.Error)

	// Free allocated memory
	// Calls the Free() function in IResultSet
	FreeResult(_result ResultSet)
}

type ResultSet interface {
	GetDataInt(_s string) int32
	GetDataLong(_s string) int64
	GetDataString(_s string) string

	// Next row
	Next() bool

	// Total rows
	Count() uint64

	// Clear resultset
	Free()
}
