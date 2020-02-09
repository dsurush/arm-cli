package controllers

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/dsurush/arm-cli/manager-cli/models"
	"github.com/dsurush/arm-core/dbupdate"
	"log"
	"os"
)

func AddClientHandler(db *sql.DB) (err error){
	fmt.Println("Enter your data")

	var newClient models.CreateNewClient
	fmt.Println("Enter users name: ")
	_, err = fmt.Scan(&newClient.Name)
	if err != nil {
		return err
	}

	fmt.Println("Enter users surname: ")
	_, err = fmt.Scan(&newClient.Surname)
	if err != nil {
		return err
	}
	// TODO: Проверка на уникальность логина
	fmt.Println("Enter users login: ")
	_, err = fmt.Scan(&newClient.Login)
	if err != nil {
		return err
	}

	fmt.Println("Enter users password: ")
	_, err = fmt.Scan(&newClient.Password)
	if err != nil {
		return err
	}
	err = dbupdate.AddClient(newClient.Name, newClient.Surname, newClient.Login, newClient.Password, db)
	if err != nil {
		log.Fatalf("Ne dobavilas")
	}

	fmt.Println("Users added successfully")
	fmt.Printf("name: %s, \nsurname: %s, \nlogin: %s,\npassword: %s", newClient.Name, newClient.Surname, newClient.Login, newClient.Password)
	return nil
}

func AddATM(db *sql.DB) (err error){

	var newATM models.CreateNewATM

	fmt.Println("Enter ATMs address")
	reader := bufio.NewReader(os.Stdin)
	newATM.Address, err = reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Can't read command: %v", err)
	}

	fmt.Println("Enter true if atm is activity, else false")
	_, err = fmt.Scan(&newATM.Locked)
	if err != nil {
		log.Printf("Ошибка при вводе данных")
		return err
	}

	err = dbupdate.AddATM(newATM.Address, newATM.Locked, db)
	if err != nil {
		log.Printf("Проблема соединения с сервером %e", err)
		return err
	}

	activity := "Не активный"
	if newATM.Locked == true{
		activity = "активный"
	}
	fmt.Printf("Был добавлен АТМ по адрессу: %s\nТип активности: %s", newATM.Address, activity)
	return nil
}