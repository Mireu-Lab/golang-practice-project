package sql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// sql.DB 객체 db 생성
var DataBase, _ = sql.Open("mysql", "root:mireu1214DB!@tcp(mireu-server.iptime.org:10000)/test")
