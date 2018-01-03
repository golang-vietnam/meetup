package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db1, _ := sql.Open("mysql", "root:123@tcp(localhost:3306)/tmp")
	db1.SetMaxOpenConns(1)

	db2, _ := sql.Open("mysql", "root:123@tcp(localhost:3306)/tmp")
	db2.SetMaxOpenConns(1)

	for {
		if err := db1.Ping(); err != nil {
			fmt.Println("[Ping failed] ------ ", err)
		}

		if _, err := db2.Exec("SELECT 1"); err != nil {
			fmt.Println("[Select failed] ----- ", err)
		}

		fmt.Printf("\n\n")
		time.Sleep(10 * time.Second)
	}
}
