package sql

import "database/sql"

var DataBase, _ = sql.Open("mysql", "root:mireu1214DB!@tcp(mireu-server.iptime.org:10000)/test")
