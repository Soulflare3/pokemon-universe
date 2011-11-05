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
	"os"
	"mysql"
)

func DBQuerySelect(_query string) (result *mysql.Result, err os.Error) {
	if err = g_db.Query(_query); err != nil {
		g_logger.Printf("[ERROR] SQL error while executing query:\n\r%s\n\rError: %s\n\r", _query, err)
		return nil, err
	}
	
	result, err = g_db.UseResult()
	if err != nil {
		g_logger.Println("[ERROR] SQL error while fetching result for query:\n\r%s\n\rError: %s\n\r", _query, err)
		return nil, err
	}
	
	return result, nil
}