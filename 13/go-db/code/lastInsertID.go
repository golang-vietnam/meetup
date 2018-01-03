package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/tmp")
	if err != nil {
		panic(err)
	}

	db.Exec("DELETE FROM person")

	var lastID int64
	if r, err := db.Exec("INSERT INTO person(name, national_id) VALUES (?,?), (?,?)",
		"A", "1",
		"B", "2"); err != nil {
		panic(err)
	} else {
		lastID, _ = r.LastInsertId()
		fmt.Println("Last insert id: ", lastID)
	}
}
