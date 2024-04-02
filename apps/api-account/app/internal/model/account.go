package model

type Account struct {
	Id        string     `json:"id"`
	Customers []Customer `json:"customers"`
}
type Customer struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Document string `json:"document"`
}
