package main

import (
	"os"
	"mysql"
)

func DBQuerySelect(_query string) (result *mysql.Result, err os.Error) {
	if err = g_db.Query(_query); err != nil {
		g_logger.Println("[ERROR] SQL error while executing query:\n\r" + _query + "\n\rError: " + err)
		return nil, err
	}
	
	result, err = g_db.UseResult()
	if err != nil {
		g_logger.Println("[ERROR] SQL error while fetching result for query:\n\r" + _query + "\n\rError: " + err)
		return nil, err
	}
	
	return result, nil
}