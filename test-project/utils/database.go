package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "local.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	CreateTable(db)

	// fmt.Println(db)
	return db
}

func CreateTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		age INTEGER
	);`

	_, err := db.Exec(query)
	return err
}
