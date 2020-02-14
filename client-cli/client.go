package main

import (
	"database/sql"
	"fmt"
	"github.com/dsurush/arm-cli/client-cli/controllers"
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
	for {
		fmt.Println(unauthorizedOperations)
		fmt.Scan(&cmd)
		switch cmd {
		case "1":
			id, err := controllers.Authorize(db)
			if err != nil {
				fmt.Println("Попробуйте еще раз")
				continue
			} else {
				controllers.AuthorizedOperations(id, db)
			}
		case "3":
			err := controllers.GetATMsForClient(db)
			if err != nil {
				fmt.Printf("Ошибка выдачи списка банкоматов %s:", err)
			}
		case "q":
			os.Exit(0)
		default:
			fmt.Println("Введена неверная команда, попробуйте еще раз", unauthorizedOperations)
			continue
		}
	}
}