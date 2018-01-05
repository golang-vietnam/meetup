package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/tmp")
	if err != nil {
		panic(err)
	}

	db.Exec("DELETE FROM person")

	var tx *sql.Tx
	if tx, err = db.Begin(); err != nil {
		panic(err)
	}
	fmt.Println("Batch insert with error:", batchInsert(tx))
}

func batchInsert(tx *sql.Tx) (err error) {
	shouldAutoRollBack := true
	defer func() {
		if err != nil && shouldAutoRollBack {
			tx.Rollback()
		}
	}()

	var stmt *sql.Stmt
	if stmt, err = tx.Prepare("INSERT INTO person(name, national_id) VALUES (?, ?)"); err != nil {
		return
	}
	for i := 0; i <= 100; i++ {
		if _, err = stmt.Exec(strconv.Itoa(i), i); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		shouldAutoRollBack = false
	}
	// done batch insert

	return
}
