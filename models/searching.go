package models

import (
	"database/sql"
	"fmt"
	"log"
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

func printItems(rows *sql.Rows) error {
	var name, description string
	var price, amount float64

	for rows.Next() {
		err := rows.Scan(&name, &description, &price, &amount)
		if err != nil {
			return err
		}

		fmt.Printf("Name: %s, Description: %s, Price: %f, Amount: %f\n", name, description, price, amount)
	}

	err := rows.Err()
	if err != nil {
		log.Printf("Error while processing rows: %v\n", err)
		return err
	}

	if rows.NextResultSet() {
		log.Println("Unexpected additional result set")
		return fmt.Errorf("Unexpected additional result set")
	}

	return nil
}
