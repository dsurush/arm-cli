package controllers

import (
	"database/sql"
	"fmt"
	"github.com/dsurush/arm-core/dbupdate"
	"log"
)

func AddClientHandler(db *sql.DB) (err error){
	fmt.Println("Enter your data")
	var name string
	fmt.Println("Enter users name: ")
	_, err = fmt.Scan(&name)
	if err != nil {
		return err
	}

	var surname string
	fmt.Println("Enter users surname: ")
	_, err = fmt.Scan(&surname)
	if err != nil {
		return err
	}

	// TODO: Проверка на уникальность логина
	var login string
	fmt.Println("Enter users login: ")
	_, err = fmt.Scan(&login)
	if err != nil {
		return err
	}

	var password string
	fmt.Println("Enter users password: ")
	_, err = fmt.Scan(&password)
	if err != nil {
		return err
	}
	err = dbupdate.AddClient(name, surname, login, password, db)
	if err != nil {
		log.Fatalf("Ne dobavilas")
	}
	fmt.Println("Users added successfully ")
	return nil
}
