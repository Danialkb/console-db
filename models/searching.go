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
	defer exec.Close()
	if exec.Next() == false {
		fmt.Println("No such items!")
	} else {
		//fmt.Println(exec.Next())

		err := printItems(exec)
		if err != nil {
			return
		}
	}
	err := exec.Close()
	if err != nil {
		return
	}
	err = db.Close()
	if err != nil {
		return
	}
}

func Some_func() {
	db, _ := initDB()

	exec, _ := db.Query("SELECT * FROM item_info;")

	print_all(exec)
	exec.Close()
	db.Close()
}
