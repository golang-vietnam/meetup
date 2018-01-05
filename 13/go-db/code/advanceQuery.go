package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// metadata definition
type metadata struct {
	Job      string
	Address  string
	WalletID int64
}

func (c *metadata) Value() (val driver.Value, err error) {
	if c == nil {
		return
	}
	return json.Marshal(&c)
}

func (c *metadata) Scan(value interface{}) (err error) {
	if value == nil {
		return
	} else if v, ok := value.([]byte); ok {
		err = json.Unmarshal(v, c)
		return
	}
	return fmt.Errorf("Can not parse metadata")
} // metadata end definition

type person struct {
	Id          int
	Name        string
	National_id int64
	Metadata    *metadata
}

func main() {
	db := connectAndClearData()

	// insert dummy person without metadata
	insertDummy(db)

	// insert jon snow with metadata
	if id, err := insertJonSnow(db); err != nil {
		panic(err)
	} else {
		fmt.Println("JonSnow has id:", id)
	}

	if persons, err := selectAllPeople(db); err != nil {
		panic(err)
	} else {
		for _, v := range persons {
			fmt.Printf("%v %+v\n", v, v.Metadata)
		}
	}
}

func insertJonSnow(db *sql.DB) (int64, error) {
	r, err := db.Exec("INSERT INTO person(name, national_id, metadata) VALUES (?,?,?)",
		"Jon Snow", 1, &metadata{
			Job:      "Actor",
			Address:  "Vietnam",
			WalletID: 1234567,
		})
	if err != nil {
		panic(err)
	}
	return r.LastInsertId()
}

func insertDummy(db *sql.DB) (int64, error) {
	r, err := db.Exec("INSERT INTO person(name, national_id) VALUES (?,?)", "Dummy", 2)
	if err != nil {
		panic(err)
	}
	return r.LastInsertId()
}

func selectAllPeople(db *sql.DB) (persons []*person, err error) {
	rows, err := db.Query("SELECT * FROM person")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // very important

	persons = make([]*person, 0, 1000)
	for rows.Next() {
		var tmp person
		rows.Scan(&tmp.Id, &tmp.Name, &tmp.National_id, &tmp.Metadata)
		persons = append(persons, &tmp)
	}

	return
}

func connectAndClearData() *sql.DB {
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/tmp")
	if err != nil {
		panic(err)
	}

	if _, err := db.Exec("DELETE FROM person"); err != nil {
		panic(err)
	}

	return db
}
