package main

import (
	"database/sql"
	"fmt"
	"github.com/dsurush/arm-core/dbinit"
	_ "github.com/mattn/go-sqlite3"
	"log"
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
	err = dbinit.Init(db)
	if err != nil {
		log.Fatal("All go with vagine")
	}

}
