package models

import (
	"fmt"
)

type Search interface {
	SearchByName(string)
}

type SearchService struct{}

func (s *SearchService) SearchByName(pattern string) {
	db, _ := initDB()

	pattern += "%"
	command := fmt.Sprintf("SELECT name, description, price, amount FROM item_info WHERE name like '%s';", pattern)
	exec, _ := db.Query(command)

	err := printItems(exec)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	exec.Close()
	db.Close()
}
