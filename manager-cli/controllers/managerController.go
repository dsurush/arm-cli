package controllers

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/dsurush/arm-cli/manager-cli/models"
	"github.com/dsurush/arm-core/dbupdate"
	"github.com/dsurush/arm-core/dbupdate/cmodels"
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
	//dbupdate.Test()
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

	fmt.Println("Введите количество денег: ")
	var balance int64
	_, err = fmt.Scan(&balance)
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
	err = dbupdate.AddAccount(clientID, paymentSystem, balance, locked, db)
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
	clientsInSlice, err := dbupdate.GetAllClients(db)
	clients := cmodels.ClientList{clientsInSlice}
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
	AccountsInSLice, err := dbupdate.GetAllAccounts(db)
	Accounts := cmodels.AccountList{AccountsInSLice}
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
	ATMsInSlice, err := dbupdate.GetAllATMs(db)
	ATMs := cmodels.AtmList{ATMsInSlice}
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
	return nil
}

func AddAtmFromXmlJson(db *sql.DB) (err error) {
	/////XML
	file, err := ioutil.ReadFile("ATM.xml")
	if err != nil {
		log.Fatalf("Can't read file %e", err)
		return err
	}
	var Atms cmodels.AtmList
	err = xml.Unmarshal(file, &Atms)
	if err != nil {
		log.Fatal("Can't Unmarshal file", err)
		return err
	}
	for _, Atm := range Atms.ATMs{
		Address := Atm.Name
		Locked := Atm.Locked
		err = dbupdate.AddATM(Address, Locked, db)
		if err != nil {
			log.Printf("Проблема соединения с сервером %e", err)
			return err
		}
	}

	////// JSON
	file, err = ioutil.ReadFile("ATM.json")
	if err != nil {
		log.Fatalf("Can't read file %e", err)
		return err
	}
	err = json.Unmarshal(file, &Atms)
	if err != nil {
		log.Fatal("Can't Unmarshal file: ", err)
		return err
	}
	for _, Atm := range Atms.ATMs{
		Address := Atm.Name
		Locked := Atm.Locked
		err = dbupdate.AddATM(Address, Locked, db)
		if err != nil {
			log.Printf("Проблема соединения с сервером %e", err)
			return err
		}
	}
	return nil
}

func AddClientsFromXmlJson(db *sql.DB) (err error){
	file, err := ioutil.ReadFile("clients.xml")
	var clients cmodels.ClientList
	err = xml.Unmarshal(file, &clients)
	if err != nil {
		log.Fatalf("Всё очень плохо, не получилось анмаршилить из клиент ксмл", err)
		return err
	}

	for _, user := range clients.Clients{
		err = dbupdate.AddClient(user.Name, user.Surname, user.Login, user.Password, user.NumberPhone, db)
		if err != nil {
			log.Fatalf("Ne tut to bilo delo")
			return err
		}
	}
	////JSON
	file, err = ioutil.ReadFile("clients.json")
	if err != nil {
		log.Fatalf("Can't read file %e", err)
		return err
	}
	var clientList cmodels.ClientList
	err = json.Unmarshal(file, &clientList)
	if err != nil {
		log.Fatal("Can't Unmarshal file: ", err)
		return err
	}
	for _, user := range clientList.Clients{
		fmt.Println(user.Name, user.Surname, user.Login, user.Password, user.NumberPhone)
		err = dbupdate.AddClient(user.Name, user.Surname, user.Login, user.Password, user.NumberPhone, db)
		if err != nil {
			log.Fatalf("Ne tut to bilo delo")
			return err
		}
	}
	return nil
}

func AddAccountsFromXmlJson(db *sql.DB) (err error) {
	file, err := ioutil.ReadFile("account.xml")
	if err != nil {
		log.Fatalf("Wring BLA %s", err)
		return err
	}
	var AccountList cmodels.AccountList
	err = xml.Unmarshal(file, &AccountList)
	if err != nil {
		log.Fatalf("Owibka BLA : %s", err)
		return err
	}

	for _, Account := range AccountList.AccountWithUserName{
		fmt.Println(Account.Client.Name, Account.Client.Surname, Account.Client.Login, Account.Client.Password, Account.Client.NumberPhone)
		err = dbupdate.AddClient(Account.Client.Name, Account.Client.Surname, Account.Client.Login, Account.Client.Password, Account.Client.NumberPhone, db)
		if err != nil {
			log.Fatalf("Ne poluchilos Add Client %s", err)
			return err
		}
		err = dbupdate.AddAccount(Account.Account.UserId, Account.Account.Name, Account.Account.Balance , Account.Account.Locked, db)
		if err != nil {
			log.Fatalf("Ne poluchilos Add Account %s", err)
			return err
		}
	}
	///JSON
	file, err = ioutil.ReadFile("account.json")
	if err != nil {
		log.Fatalf("Wring BLA %s", err)
		return err
	}
//	var AccountList cmodels.AccountList
	err = json.Unmarshal(file, &AccountList)
	if err != nil {
		log.Fatalf("Owibka BLA : %s", err)
		return err
	}

	for _, Account := range AccountList.AccountWithUserName{
		fmt.Println(Account.Client.Name, Account.Client.Surname, Account.Client.Login, Account.Client.Password, Account.Client.NumberPhone)
		err = dbupdate.AddClient(Account.Client.Name, Account.Client.Surname, Account.Client.Login, Account.Client.Password, Account.Client.NumberPhone, db)
		if err != nil {
			log.Fatalf("Ne poluchilos Add Client %s", err)
			return err
		}
		err = dbupdate.AddAccount(Account.Account.UserId, Account.Account.Name, Account.Account.Balance , Account.Account.Locked, db)
		//err = dbupdate.AddAccount(Account.Account.UserId, Account.Account.Name, Account.Account.Locked, db)
		if err != nil {
			log.Fatalf("Ne poluchilos Add Account %s", err)
			return err
		}
	}
	return nil
}