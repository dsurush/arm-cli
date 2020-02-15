package main

import (
	"database/sql"
	"encoding/json"
	"github.com/dsurush/arm-cli/manager-cli/controllers"
	"github.com/dsurush/arm-core/dbinit"
	"github.com/dsurush/arm-core/dbupdate/cmodels"
	"os"
	"strings"

	//"encoding/xml"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
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
	mainAppFunction(db)
}

func mainAppFunction(db *sql.DB) {
	var cmd string
	for {
		fmt.Println(AuthorizedOperations)
		fmt.Scan(&cmd)
		switch cmd {
		case "1":
			controllers.AddClientHandler(db)
		case "2":
			controllers.AddAccountHandler(db)
		case "3":
			controllers.AddServiceHandler(db)
		case "4":
			controllers.AddClientsToJsonXmlFiles(db)
		case "5":
			controllers.AddAccountsToJsonXmlFiles(db)
		case "6":
			controllers.AddATMsToJsonXmlFiles(db)
		case "7":
			controllers.AddClientsFromXmlJson(db)
		case "8":
			controllers.AddAccountsFromXmlJson(db)
		case "9":
			controllers.AddAtmFromXmlJson(db)
		case "10":
			controllers.AddATMHandler(db)
		case "q":
			os.Exit(0)
		default:
			fmt.Println("Введенно неверное значение, пробуйте еще раз\n")
			continue
		}
	}
}

func test(db *sql.DB)  {
	file, err := ioutil.ReadFile("ATM.json")
	if err != nil {
		log.Fatal(err)
	}
	var backup cmodels.AtmList
	err = json.Unmarshal(file, &backup)
	if err != nil {
		log.Fatal("hello", err)
		//re/turn err
	}

	for _, value := range backup.ATMs {
		//String
		split := strings.Split(value.Name, "\n")
		value.Name = split[0]
		fmt.Println(value.ID, value.Name, value.Locked)
	}
}