package models

type Client struct{
	ID int
	Name string
	Surname string
	Login string
	Password string
	locked bool
	NumberPhone string
}

type CreateNewClient struct {
	Name string
	Surname string
	Login string
	Password string
	NumberPhone string
}