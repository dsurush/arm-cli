package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatalf("Ошибка открытия базы данных %s", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Нет подключения к серверу")
	}
	mainAppFunction(db)
}
func mainAppFunction(db *sql.DB)  {
	var cmd string
	fmt.Println(unauthorizedOperations)
	for {
		fmt.Scan(&cmd)
		switch cmd {
		case "1":
		case "q":
			os.Exit(0)
		default:

			fmt.Println("Введена неверная команда, попробуйте еще раз", unauthorizedOperations)
			continue
		}
	}
}
