package models

type Client struct{
	ID int
	Name string
	Surname string
	Login string
	Password string
	locked bool
}

type CreateNewClient struct {
	Name string
	Surname string
	Login string
	Password string
}