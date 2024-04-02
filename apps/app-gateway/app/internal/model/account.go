package model

type Account struct {
	Id        string     `json:"id"`
	Customers []Customer `json:"customers"`
}
