package main

import (
	"fmt"
	"mysql"
	"time"
)

const (
	MESSAGE_ACCEPTED = "Your map submission has been accepted."
	MESSAGE_DECLINED = "Your map sumission has been declined."
)

func InjectMessage(_username string, _message string, _accepted bool) {
	dbconn, err := mysql.DialTCP("94.75.231.83", "extern", "supersecret9001", "pu-web")
	if err != nil {
		fmt.Printf("Could not connect to website database: %s\n", err)
		return
	}

	userQuery := fmt.Sprintf("SELECT id_member FROM smf_members WHERE member_name='%s'", _username)
	if err := dbconn.Query(userQuery); err != nil {
		fmt.Printf("Query error: %s\n", err)
		return
	}

	result, err := dbconn.UseResult()
	if err != nil {
		fmt.Printf("Query error #2: %s", err)
		return
	}
	row := result.FetchMap()
	if row != nil {
		id_member := row["id_member"].(uint)

		result.Free()

		userQuery := fmt.Sprintf("UPDATE smf_members SET unread_messages = unread_messages + 1 WHERE id_member='%d'", id_member)
		if err := dbconn.Query(userQuery); err != nil {
			fmt.Printf("Query error #3: %s\n", err)
			return
		}

		message := MESSAGE_ACCEPTED
		if !_accepted {
			message = MESSAGE_DECLINED
		}
		if len(_message) > 0 {
			message = fmt.Sprintf("%s\n\n%s", message, _message)
		}

		pmQuery := fmt.Sprintf("INSERT INTO smf_personal_messages (id_member_from, from_name, msgtime, subject, body) VALUES ('0', 'Pokemon Universe MMORPG', '%d', 'Map Submission', '%s')", time.Now(), message)
		if err := dbconn.Query(pmQuery); err != nil {
			fmt.Printf("Database query error: %s\n", err)
			return
		}
		pmid := dbconn.LastInsertId

		messageQuery := fmt.Sprintf("INSERT INTO smf_pm_recipients (id_pm, id_member, labels, is_read, is_new) VALUES ('%d', '%d', '-1', '0', '1')", pmid, id_member)
		if err := dbconn.Query(messageQuery); err != nil {
			fmt.Printf("Database query error #2: %s\n", err)
			return
		}
	} else {
		result.Free()
	}
}
