package sql

import (
	"database/sql"
	"log"
)

func SettingTestTable() *sql.DB {
	myDB, err_open := sql.Open("sqlite3", "./sql/sqlite.db")

	if err_open != nil {
		log.Print(err_open)
	}

	createTableQuery := `
		CREATE TABLE IF NOT EXISTS test_table (
			id INTEGER NOT NULL PRIMARY KEY,
			name TEXT,
			age TEXT,
			UNIQUE (name)
		);
	`

	_, err_exec := myDB.Exec(createTableQuery)
	if err_exec != nil {
		log.Print(err_exec)
	}

	return myDB
}
