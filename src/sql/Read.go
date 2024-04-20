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

// Read function reads data from the database and encodes it as JSON
func Read() interface{} {
	// Query the database and get the rows
	var rows, _ = DataBase.Query("SELECT * FROM `untitled_table`")

	var TableData TableInfo
	var msg []TableInfo

	// Iterate over the rows and scan the columns into the TableData struct
	for rows.Next() {
		rows.Scan(
			&TableData.ID,
			&TableData.Message,
			&TableData.Time)

		// Append the TableData struct to the msg slice
		msg = append(msg, TableData)
	}

	// Encode the msg slice as JSON and write it to stdout
	return json.NewEncoder(os.Stdout).Encode(msg)
}
