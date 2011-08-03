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

//#include <stdlib.h>
//#include <mysql/mysql.h>
//#include <mysql/errmsg.h>
//
//char *wm_row(MYSQL_ROW row, int i) {
//    return (char *)(row)[i];
//}
import "C"
import (
	"os"
	"unsafe"
	"fmt"
	"strconv"
)

type MySQL struct {
	handle      C.MYSQL
	isConnected bool
}

func NewMySQL() *MySQL {
	return &MySQL{isConnected: false}
}

func (self *MySQL) getLastError(_message string) os.Error {
	return os.NewError(fmt.Sprintf("%v. MYSQL_ERROR: %v", _message, C.mysql_error(&self.handle)))
}

func (self *MySQL) getLastErrorNo() int {
	return int(C.mysql_errno(&self.handle))
}

func (self *MySQL) Connect(_client ClientInfo) (err os.Error) {
	if C.mysql_init(&self.handle) == nil {
		err = self.getLastError("Failed to initialize MySQL connection handle.")
		return
	}

	var reconnect C.my_bool = 1
	C.mysql_options(&self.handle, C.MYSQL_OPT_RECONNECT, unsafe.Pointer(&reconnect))

	host := C.CString(_client.SQLHost)
	user := C.CString(_client.SQLUser)
	pass := C.CString(_client.SQLPass)
	dbna := C.CString(_client.SQLDB)
	ret := C.mysql_real_connect(&self.handle, host, user, pass, dbna, 3306, nil, 0)
	if ret == nil {
		err = self.getLastError("Failed to connect to database.")
		return
	}

	if C.MYSQL_VERSION_ID < 50019 {
		//mySQL servers < 5.0.19 has a bug where MYSQL_OPT_RECONNECT is (incorrectly) reset by mysql_real_connect calls
		//See http://dev.mysql.com/doc/refman/5.0/en/mysql-options.html for more information.
		C.mysql_options(&self.handle, C.MYSQL_OPT_RECONNECT, unsafe.Pointer(&reconnect))
		fmt.Println("Outdated mySQL server detected. Consider upgrading to a newer version.")
	}

	self.isConnected = true

	return
}

func (self *MySQL) Close() {
	if self.isConnected {
		C.mysql_close(&self.handle)
	}
}

func (self *MySQL) ExecuteQuery(_query string) (err os.Error) {
	if !self.isConnected {
		err = os.NewError("Failed to execute query (" + _query + "). Not connected.")
		return
	}

	// Execute the query
	res := C.mysql_real_query(&self.handle, C.CString(_query), C.ulong(len(_query)))
	if res != 0 {
		err = self.getLastError("Execute query: " + _query)
		error := self.getLastErrorNo()
		if error == C.CR_SERVER_LOST || error == C.CR_SERVER_GONE_ERROR {
			self.isConnected = false
		}
	}

	// we should call that every time as someone would call ExecuteQuery('SELECT...')
	// as it is described in MySQL manual: "it doesn't hurt" :P
	m_res := C.mysql_store_result(&self.handle)
	if m_res != nil {
		C.mysql_free_result(m_res)
	}

	return
}

func (self *MySQL) StoreQuery(_query string) (res ResultSet, err os.Error) {
	if !self.isConnected {
		err = os.NewError("Failed to execute query (" + _query + "). Not connected.")
		return
	}

	// Execute the query
	ret := C.mysql_real_query(&self.handle, C.CString(_query), C.ulong(len(_query)))
	if ret != 0 {
		err = self.getLastError("Execute query: " + _query)
		error := self.getLastErrorNo()
		if error == C.CR_SERVER_LOST || error == C.CR_SERVER_GONE_ERROR {
			self.isConnected = false
		}
	}

	// we should call that every time as someone would call executeQuery('SELECT...')
	// as it is described in MySQL manual: "it doesn't hurt" :P
	m_res := C.mysql_store_result(&self.handle)

	// Error occured
	if m_res == nil {
		err = self.getLastError("Store query: " + _query)
		error := self.getLastErrorNo()
		if error == C.CR_SERVER_LOST || error == C.CR_SERVER_GONE_ERROR {
			self.isConnected = false
		}

		return
	}

	res = NewMySQLResult(m_res)

	return
}

func (self *MySQL) FreeResult(_result ResultSet) {
	_result.Free()
}

/********************************
******** MYSQL DB RESULT ********
********************************/
type ResultNames map[string]uint32
type MySQLResult struct {
	listNames ResultNames
	handle    *C.MYSQL_RES
	row       C.MYSQL_ROW
}

func NewMySQLResult(_handle *C.MYSQL_RES) (res *MySQLResult) {
	res = &MySQLResult{handle: _handle}
	res.listNames = make(ResultNames)

	i := uint32(0)
	var field *C.MYSQL_FIELD
	for {
		field = C.mysql_fetch_field(_handle)
		if field == nil {
			break
		}
		res.listNames[C.GoString(field.name)] = i
		i++
	}

	return
}

func (res *MySQLResult) GetDataInt(_s string) int32 {
	id, found := res.listNames[_s]
	if found {
		value := C.wm_row(res.row, C.int(id))
		if value != nil {
			return int32(C.atoi(value))
		}
	}

	return 0
}

func (res *MySQLResult) GetDataLong(_s string) int64 {
	id, found := res.listNames[_s]
	if found {
		value := C.wm_row(res.row, C.int(id))
		if value != nil {
			ret, _ := strconv.Atoi64(C.GoString(value))
			return ret
		}
	}

	return 0
}

func (res *MySQLResult) GetDataString(_s string) string {
	id, found := res.listNames[_s]
	if found {
		value := C.wm_row(res.row, C.int(id))
		if value != nil {
			return C.GoString(value)
		}
	}

	return ""
}

func (res *MySQLResult) Next() bool {
	res.row = C.mysql_fetch_row(res.handle)
	return (res.row != nil)
}

func (res *MySQLResult) Count() uint64 {
	return uint64(C.mysql_num_rows(res.handle))
}

func (res *MySQLResult) Free() {
	C.mysql_free_result(res.handle)
}
