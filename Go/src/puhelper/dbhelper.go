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
package puhelper

import (
	"gomysql"
	"nonamelib/log"
)

var (
	DBCon	*mysql.Client
)

func DBQuerySelect(_query string) (result *mysql.Result, err error) {
	// Lock
	DBLock()
	
	err = DBCon.Query(_query)
	if err != nil {
		log.Println("[ERROR] SQL error while executing query:")
		log.Printf("%s\n", _query)
		log.Printf("Error: %s\n", err.Error()) 
		result = nil
	} else {
		result, err = DBCon.UseResult()
		if err != nil {
			log.Println("[ERROR] SQL error while fetching result for query:\n\r%s\n\rError: %s", _query, err.Error())
			result = nil
		}
	}
	
	// Unlock
	if err != nil {
		DBUnlock()
	}

	return
}

func DBQuery(_query string) (err error) {
	// Lock
	DBLock()
	defer DBUnlock()
	
	if err = DBCon.Query(_query); err != nil {
		log.Println("[ERROR] SQL error while executing query:")
		log.Printf("%s\n", _query)
		log.Printf("Error: %s\n", err.Error()) 
	}
	
	return
}

func DBQueryNoLock(_query string) (err error) {
	if err = DBCon.Query(_query); err != nil {
		log.Println("[ERROR] SQL error while executing query:")
		log.Printf("%s\n", _query)
		log.Printf("Error: %s\n", err.Error()) 
	}
	
	return
}

func DBFree() {
	DBCon.FreeResult()
	DBCon.Unlock()
}

func DBLock() {
	DBCon.Lock()
}

func DBUnlock() {
	DBCon.Unlock()
}

func DBStartTransaction() {
	DBLock()
	if err := DBCon.Start(); err != nil {	
		log.Println("[ERROR] SQL error while starting a new transaction.")
		log.Printf("Error: %s\n", err.Error())
		
		DBUnlock() 
	}
}

func DBCommit() {
	if err := DBCon.Commit(); err != nil {
		log.Println("[ERROR] SQL error while committing transaction.")
		log.Printf("Error: %s\n", err.Error()) 
		
		DBRollback()
	} else {
		DBUnlock()
	}
}

func DBRollback() {
	if err := DBCon.Rollback(); err != nil {
		log.Println("[ERROR] SQL error while rolling transaction back")
		log.Printf("Error: %s\n", err.Error())
	}
	
	DBUnlock()
}

func DBGetLastInsertId() uint64 {
	return DBCon.LastInsertId
}

func DBGetString(_row interface{}) string {
	if _row != nil {
		return _row.(string)
	}
	return ""
}

func DBGetStringFromArray(_row interface{}) string {
	if _row != nil {
		return string(_row.([]uint8))
	}
	return ""
}

func DBGetInt(_row interface{}) int {
	if _row != nil {
		return int(_row.(int64))
	}
	return 0
}

func DBGetInt64(_row interface{}) int64 {
		if _row != nil {
		return _row.(int64)
	}
	return 0
}

func DBGetUint64(_row interface{}) uint64 {
		if _row != nil {
		return _row.(uint64)
	}
	return 0
}

func DBGetFloat64(_row interface{}) float64 {
	if _row != nil {
		return float64(_row.(int64))
	}
	return 0
}

func Escape(_data string) string {
	return DBCon.Escape(_data)
}