package main

import (
	"fmt"
	"mysql"
)

func DBQuerySelect(_query string) (result *mysql.Result, err error) {
	if err = g_db.Query(_query); err != nil {
		fmt.Printf("[ERROR] SQL error while executing query:\n\r%s\n\rError: %s", _query, err)
		return nil, err
	}

	result, err = g_db.UseResult()
	if err != nil {
		fmt.Println("[ERROR] SQL error while fetching result for query:\n\r%s\n\rError: %s", _query, err)
		return nil, err
	}

	return result, nil
}

func DBQuery(_query string) (err error) {
	if err := g_db.Query(_query); err != nil {
		fmt.Printf("[ERROR] SQL error while executing query:\n\r%s\n\rError: %s", _query, err)
		return err
	}
	
	return nil
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