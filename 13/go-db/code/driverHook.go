package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// DDL
const (
	dropTable   = `DROP TABLE IF EXISTS person`
	createTable = `
CREATE TABLE person (
	id int(11) NOT NULL AUTO_INCREMENT,
	name varchar(20) DEFAULT NULL,
	national_id long NOT NULL DEFAULT 0,
	metadata longblob DEFAULT NULL,
	PRIMARY KEY (id)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
`
) // END DDL

func main() {
	// configure dsn
	dsn := "root:123@tcp(localhost:3306)/tmp"

	// open connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Failed to create *DB object", err)
	} else {
		fmt.Println("Success to create *DB object")
	}

	// run script
	if err = script(db); err != nil {
		panic(err)
	}

	fmt.Println("Script ran well!")
}

func script(db *sql.DB) (err error) {
	if err = prepare(db); err != nil {
		return
	}

	return
}

func prepare(db *sql.DB) (err error) {
	if _, err = db.Exec(dropTable); err != nil {
		return
	}

	if _, err = db.Exec(createTable); err != nil {
		return
	}

	return
}
