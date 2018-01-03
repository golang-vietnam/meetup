package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/tmp")
	if err != nil {
		panic(err)
	}

	var lastID int64
	if r, err := db.Exec("INSERT INTO person(name, national_id) VALUES (?,?)", "A", "1"); err != nil {
		panic(err)
	} else {
		lastID, _ = r.LastInsertId()
		fmt.Println("Last insert id: ", lastID)
	}

	if _, err := db.ExecContext(context.Background(),
		"DELETE FROM person WHERE id=@myID",
		sql.Named("myID", lastID)); err != nil {
		fmt.Println("You are right!", err)
	}
}
