package sql

import (
	"encoding/json"
	"os"
)

type TableInfo struct {
	ID      int    `json:"id"`
	Message string `json:"msg"`
	Time    string `json:"time"`
}

func Read() {
	var rows, _ = DataBase.Query("SELECT * FROM `untitled_table`")

	var TableData TableInfo
	var msg []TableInfo

	for rows.Next() {
		rows.Scan(
			&TableData.ID,
			&TableData.Message,
			&TableData.Time)

		msg = append(msg, TableData)
	}

	return json.NewEncoder(os.Stdout).Encode(msg)
}
