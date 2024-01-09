package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var err error

func InitDB() {
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("could not open database")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()

	fmt.Println("connected to sqlite database!")
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    name TEXT NOT NULL,
	    description TEXT NOT NULL,
	    location TEXT NOT NULL,
	    dateTime DATETIME NOT NULL,
	    user_id INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("could not create events tables")
	}

}
