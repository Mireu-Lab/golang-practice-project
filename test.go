package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// sql.DB 객체 db 생성
	var db, err1 = sql.Open("mysql", "root:mireu1214DB!@tcp(mireu-server.iptime.org:10000)/test")
	if err1 != nil {
		panic(err1)
	}

	// db 차후에 닫기
	defer db.Close()

	// SELECT 쿼리
	var rows, err2 = db.Query("SELECT * FROM `untitled_table`")
	if err2 != nil {
		panic(err2)
	}

	for rows.Next() {
		var id int
		var msg string
		var time string

		rows.Scan(&id, &msg, &time)

		fmt.Printf("ID: %d, 메시지: %s, 시간: %s\n", id, msg, time)
	}
}
