package sql

import (
	"fmt"
)

func Write(msg string) int {
	// INSERT 실행
	result, err := DataBase.Exec(`insert into untitled_table (msg, created_at) values (?, DEFAULT);`, msg)
	if err != nil {
		fmt.Println(err)
		return 500
	}

	nRow, err := result.RowsAffected()
	fmt.Println(nRow)

	if err == nil {
		return 200
	} else {
		fmt.Println(err)
		return 500
	}
}
