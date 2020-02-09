package models

type ATM struct {
	id int
	address string
	locked bool
}

type CreateNewATM struct {
	Address string
	Locked bool
}