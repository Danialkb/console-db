package models

import (
	"fmt"
	"log"
)

type Filter interface {
	FilterByPrice(sortForm string)
	FilterByRating(sortForm string)
}

type FilterService struct{}

func (f *FilterService) FilterByPrice(sortForm string) {
	db, _ := initDB()
	command := fmt.Sprintf("SELECT name, description, price, amount FROM item_info ORDER BY price %s", sortForm)
	exec, err := db.Query(command)
	if err != nil {
		log.Fatalf("Error ", err)
	}

	err = printItems(exec)
	if err != nil {
		log.Fatalf("Error ", err)
	}
}

