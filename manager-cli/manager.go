package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dsurush/arm-cli/manager-cli/models"
	"github.com/dsurush/arm-core/dbinit"
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
	//clients, err := dbupdate.GetAllClients(db)
	//if/ err != nil {
	//	fmt.Println("wrong ans\n")
	//}
	//for _, client := range clients{
	//	fmt.Println(client)
	//}
	//controllers.AddAccountHandler(db)
	//for i := 1; i <= 5; i++ {controllers.AddClientHandler(db)}
	//controllers.AddATMHandler(db)
	//controllers.AddServiceHandler(db)
	//controllers.AddClientsToJsonXmlFiles(db)
	//controllers.AddAccountsToJsonXmlFiles(db)
		test(db)
	//controllers.AddATMsToJsonXmlFiles(db)
}

func mainAppFunction() {
	var cmd string
	for {
		fmt.Scan(&cmd)
	}
}

func test(db *sql.DB)  {
	//client := models.Client{
	//	ID:          0,
	//	Name:        "1",
	//	Surname:     "1",
	//	NumberPhone: "2",
	//	Login:       "2",
	//	Password:    "2",
	//	Locked:      false,
	//}
	file, err := ioutil.ReadFile("ATM.json")
	if err != nil {
		log.Fatal(err)
	}
	var backup []models.ATM
	json.Unmarshal(file, &backup)
	if err != nil {
		log.Fatal(err)
		//re/turn err
	}
	for _, value := range backup{
		fmt.Println(value)
	}
}