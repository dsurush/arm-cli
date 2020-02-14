package controllers

import (
	"database/sql"
	"fmt"
	"github.com/dsurush/arm-core/dbupdate"
)

func Authorize(db *sql.DB) (id int64, err error){
	fmt.Println("Введите логин:")
	var login string
	fmt.Scan(&login)
	fmt.Println("Введите пароль:")
	var password string
	fmt.Scan(&password)

	predicate, err := dbupdate.Login(login, password, db)
	if predicate == false{
		fmt.Println("Введен неправильный логин")
		return 0, err
	}

	if predicate == true && err != nil{
		fmt.Println("Введен неправильный пароль")
		return 0, err
	}
	id, surname := dbupdate.SearchByLogin(login, db)

	fmt.Printf("Добро пожаловать мистер %s\n", surname)
	return id, nil
}

func GetATMsForClient(db *sql.DB) (err error){
	ms, err := dbupdate.GetAllATMs(db)
	if err != nil {
		return err
	}
	i:=0
	for _, value := range ms{
		i++
		fmt.Println(value)
	}
	if i == 0{
		fmt.Println("Список банкоматов пуст")
	}
	return nil
}

func SearchAccountByIdHandler(id int64, db *sql.DB) (err error){
	list, err := dbupdate.SearchAccountById(id, db)
	if err != nil {
		fmt.Errorf("cant : %e", err)
		return err
	}
	fmt.Println("Список ваших счетов:")
	for index, account := range list{
		fmt.Println(index + 1, ".", account.Name, account.AccountNumber, account.Locked)
	}
	return nil
}

func AuthorizedOperations(id int64, db *sql.DB){
	var cmd string
	for {
		fmt.Println(AuthorizedTextOperations)
		fmt.Scan(&cmd)
		switch cmd {
			case "1":
				SearchAccountByIdHandler(id, db)
				//TODO: список счетов
			case "2":
				//TODO: перевод денег
			case "3":
				//TODO: Оплачивать услугу
			case "q":
				return
		}
	}
}