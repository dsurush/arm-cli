package controllers

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/dsurush/arm-cli/manager-cli/models"
	"github.com/dsurush/arm-core/dbupdate"
	"io/ioutil"
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

	fmt.Println("Enter users phone: ")
	_, err = fmt.Scan(&newClient.NumberPhone)
	if err != nil {
		return err
	}

	err = dbupdate.AddClient(newClient.Name, newClient.Surname, newClient.Login, newClient.Password, newClient.NumberPhone, db)
	if err != nil {
		log.Fatalf("Ne dobavilas")
	}

	fmt.Println("Users added successfully")
	fmt.Printf("name: %s,\nsurname: %s,\nlogin: %s,\npassword: %s,\nphoneNumber: %s", newClient.Name, newClient.Surname, newClient.Login, newClient.Password, newClient.NumberPhone)
	return nil
}

func AddATMHandler(db *sql.DB) (err error){

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
	dbupdate.Test()
	return nil
}

func AddAccountHandler(db *sql.DB) (err error){
	fmt.Println("Введите ID пользователя: ")
	var clientID int64
	_, err = fmt.Scan(&clientID)
	if err != nil {
		return err
	}

	fmt.Println("Введите название типа платежной системы: ")
	var paymentSystem string
	_, err = fmt.Scan(&paymentSystem)
	if err != nil {
		return err
	}
	fmt.Println("Введите 1 если хотите разблокировать сейчас же счет, иначе 0:")
	locked := false
	var typeOfLock int
	_, err = fmt.Scan(&typeOfLock)
	if err != nil {
		return err
	}
	if typeOfLock == 1{
		locked = true
	}
	err = dbupdate.AddAccount(clientID, paymentSystem, locked, db)
	if err != nil {
		fmt.Errorf("Ошибка при добавлении, %e", err)
	}
	return nil
}

func AddServiceHandler(db *sql.DB) (err error) {
	fmt.Println("Введите название услуги:")
	var serviceName string
	reader := bufio.NewReader(os.Stdin)
	serviceName, err = reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Can't read command: %v", err)
		return err
	}

	fmt.Println("Введите цену услуги: ")
	var price int64
	_, err = fmt.Scan(&price)
	if err != nil {
		fmt.Errorf("Wrongerr %e", err)
		return err
	}

	err = dbupdate.AddService(serviceName, price, db)
	if err != nil {
		fmt.Errorf("errorr %e", err)
		return err
	}
	return nil
}

func AddClientsToJsonXmlFiles(db *sql.DB) (err error){
	clients, err := dbupdate.GetAllClients(db)
	if err != nil {
		fmt.Errorf("Ошибка при получении клиентов с БД %e", err)
		return err
	}
	////Json
	data, err := json.Marshal(clients)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = ioutil.WriteFile("clients.json", data, 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}
	////XML
	data, err = xml.Marshal(clients)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = ioutil.WriteFile("clients.xml", data, 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}
	////
	return nil
}

func AddAccountsToJsonXmlFiles(db *sql.DB) (err error){
	Accounts, err := dbupdate.GetAllAccounts(db)
	if err != nil {
		fmt.Errorf("Ошибка при получении клиентов с БД %e", err)
		return err
	}
	////Json
	data, err := json.Marshal(Accounts)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = ioutil.WriteFile("account.json", data, 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}
	////XML
	data, err = xml.Marshal(Accounts)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = ioutil.WriteFile("account.xml", data, 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}
	////
	return nil
}

func AddATMsToJsonXmlFiles(db *sql.DB) (err error){
	ATMs, err := dbupdate.GetAllATMs(db)
	if err != nil {
		fmt.Errorf("Ошибка при получении клиентов с БД %e", err)
		return err
	}
	////Json
	data, err := json.Marshal(ATMs)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = ioutil.WriteFile("ATM.json", data, 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}
	////XML
	data, err = xml.Marshal(ATMs)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = ioutil.WriteFile("ATM.xml", data, 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}
	////
	return nil
}
