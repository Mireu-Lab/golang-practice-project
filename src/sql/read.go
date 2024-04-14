package sql

import "fmt"

func Read() {
	var name string

	err := DataBase.QueryRow("SELECT * FROM `untitled_table`;").Scan(&name)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(name)
}
