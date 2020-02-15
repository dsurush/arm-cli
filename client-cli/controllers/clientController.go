package controllers

import (
	"database/sql"
	"fmt"
	"github.com/dsurush/arm-core/dbupdate"
	"log"
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

func SearchAccountByIdHandler(id int64, db *sql.DB) (accounts map[int64]int64, err error){
	list, err := dbupdate.SearchAccountById(id, db)
	accounts = map[int64]int64{}
	if err != nil {
		fmt.Errorf("cant : %e", err)
		return nil, err
	}
	fmt.Println("Список ваших счетов:")
//	var index64 int64
	for index, account := range list {
		index64 := int64(1 + index)
		accounts[index64] = account.AccountNumber
		fmt.Println(index+1, ".", account.Name, account.AccountNumber, account.Balance)
	}
	return accounts, nil
}
/*
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
*/
func AuthorizedOperations(id int64, db *sql.DB){
	var cmd string
	for {
		fmt.Println(AuthorizedTextOperations)
		fmt.Scan(&cmd)
		switch cmd {
			case "1":
				SearchAccountByIdHandler(id, db)
			case "2":
				//TODO: перевод денег
			case "3":
				err := PayServiceHandler(id, db)
				if err != nil {
					log.Fatal("Uliya")
				}
				//TODO: Оплачивать услугу
			case "q":
				return
		}
	}
}


func PayServiceHandler(id int64, db *sql.DB) (err error){
	//TODO: Вывести список счетов, где можно выбрать из списка счетов АккаунтНамбер
	fmt.Println("Выберите счет:")
	accounts, err := SearchAccountByIdHandler(id, db)
	if err != nil {
		return err
	}
	fmt.Println(accounts)

	for {
		var cmd int64
		fmt.Scan(&cmd)
		switch (int64(len(accounts)) >= cmd && cmd > 0 ) {
		case true:
			//TODO Выберите продукт
			fmt.Println(accounts[cmd])
			ChooseToService(accounts[cmd], db)
			return nil
		case false:
			fmt.Println("Введите заново в пределах количество ваших счетов")
		}
	}
	return nil
}

func GetAllServicesHandler(db *sql.DB) (err error){
	services, err := dbupdate.GetAllServices(db)
	if err != nil {
		fmt.Errorf("Get all services didn't work %e", err)
		return nil
	}

	for _, service := range services{
		fmt.Println(service.ID, service.Name, service.Price)
	}
	return nil
}
type NewErr struct {
	err error
}
func ChooseToService(AccountNumber int64, db *sql.DB) (err error){
	fmt.Println("Выберите услугу: ")
	err = GetAllServicesHandler(db)
	if err != nil {
		fmt.Errorf("GetServiceHandler %e", err)
		return err
	}
	for {
		var cmd int64
		fmt.Scan(&cmd)
		err := dbupdate.CheckServiceHaving(cmd, db)
		if err != nil{
			fmt.Println("Такой услуги нет, попробуйте еще раз")
			continue
		} else {
			err := Transfer(AccountNumber, cmd, db)
			if err != nil {
				fmt.Println("Перевод невозможен")
			}
			//TODO Tranz (accountNumber --> cmd (service ID)
		}
		return nil
	}
}

func Transfer(accountNumber, ServiceID int64, db *sql.DB) (err error){
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	var AccountBalance int64
	err = tx.QueryRow(`select balance from accounts where accountNumber = ?`, accountNumber).Scan(&AccountBalance)
	if err != nil {
		return err
	}
	var ServicePrice int64
	err = tx.QueryRow(`select price from services where id = ?`, ServiceID).Scan(&ServicePrice)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`UPDATE services set balance = balance + price where id = ?`, ServiceID)
	if err != nil {
		return err
	}
	_, err = tx.Exec(`UPDATE accounts set balance = balance - ? where accountNumber = ?`, ServicePrice, accountNumber)
	if err != nil {
		return err
	}

	return nil
}