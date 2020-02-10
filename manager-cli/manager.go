package main

import (
	"database/sql"
	"fmt"
	"github.com/dsurush/arm-cli/manager-cli/controllers"
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
		fmt.Println("Нет подключения к серверу")
	}

	err = dbinit.Init(db)
	if err != nil {
		log.Fatal("All go with vagine")
	}
	//controllers.AddAccountHandler(db)
	//controllers.AddClientHandler(db)
	//controllers.AddATMHandler(db)
	controllers.AddServiceHandler(db)
}

func mainAppFunction() {
	var cmd string
	for {
		fmt.Scan(&cmd)
	}
}