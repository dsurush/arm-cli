package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)
func main() {
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		fmt.Println("Owibka", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("ALO")
	}

	fmt.Println("Hello")
}
