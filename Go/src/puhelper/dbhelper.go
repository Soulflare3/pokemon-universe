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
	"putools/log"
)

var (
	DBCon	*mysql.Client
)

func DBQuerySelect(_query string) (result *mysql.Result, err error) {
	if err = DBCon.Query(_query); err != nil {
		logger.Println("[ERROR] SQL error while executing query:")
		logger.Printf("%s\n", _query)
		logger.Printf("Error: %s\n", err.Error()) 
		return nil, err
	}

	result, err = DBCon.UseResult()
	if err != nil {
		logger.Println("[ERROR] SQL error while fetching result for query:\n\r%s\n\rError: %s", _query, err.Error())
		return nil, err
	}

	return result, nil
}

func DBQuery(_query string) (err error) {
	if err := DBCon.Query(_query); err != nil {
		logger.Println("[ERROR] SQL error while executing query:")
		logger.Printf("%s\n", _query)
		logger.Printf("Error: %s\n", err.Error()) 
		return err
	}
	
	return nil
}

func DBStartTransaction() {
	if err := DBCon.Start(); err != nil {
		logger.Println("[ERROR] SQL error while starting a new transaction.")
		logger.Printf("Error: %s\n", err.Error()) 
	}
}

func DBCommit() {
	if err := DBCon.Commit(); err != nil {
		logger.Println("[ERROR] SQL error while committing transaction.")
		logger.Printf("Error: %s\n", err.Error()) 
		
		DBRollback()
	}
}

func DBRollback() {
	if err := DBCon.Rollback(); err != nil {
		logger.Println("[ERROR] SQL error while rolling transaction back")
		logger.Printf("Error: %s\n", err.Error())
	}
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

func DBGetFloat64(_row interface{}) float64 {
	if _row != nil {
		return float64(_row.(int64))
	}
	return 0
}

func Escape(_data string) string {
	return DBCon.Escape(_data)
}