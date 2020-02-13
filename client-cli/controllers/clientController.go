package controllers

import (
	"database/sql"
	"fmt"
	"github.com/dsurush/arm-core/dbupdate"
)

func Authorize(db *sql.DB) (err error){
	fmt.Println("Введите логин:")
	var login string
	fmt.Scan(&login)
	fmt.Println("Введите пароль:")
	var password string
	fmt.Scan(&password)

	predicate, err := dbupdate.Login(login, password, db)
	if predicate == false{
		fmt.Println("Введен неправильный логин")
		return err
	}

	if predicate == true && err != nil{
		fmt.Println("Введен неправильный пароль")
		return err
	}
	Surname := dbupdate.SearchByLogin(login, db)
	fmt.Printf("Добро пожаловать мистер %s\n", Surname)
	return nil
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
func AuthorizedOperations(db *sql.DB){
	var cmd string
	for {
		fmt.Println(AuthorizedTextOperations)
		fmt.Scan(&cmd)
		switch cmd {
			case "1":
				//TODO: список счетов
			case "2":
				//TODO: перевод денег
			case "3":
				//TODO: Оплачивать услугу
		}
	}
}