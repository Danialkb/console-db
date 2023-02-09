package models

type User struct {
	Login    string
	Password string
}

type Item struct {
	Name        string
	Description string
	Amount      int
	Price       float64
}
