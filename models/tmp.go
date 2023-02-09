package models

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func initDB() (*sql.DB, error) {
	db, err := sql.Open("postgres",
		"user=postgres password=Dankb2131193* host=localhost port=5432 dbname=assignment1_console_golang sslmode=disable")

	if err != nil {
		return nil, err
	}

	return db, nil
}

//func (u *User) FilterItemByPrice(sortForm string) {
//	db, _ := initDB()
//	command := fmt.Sprintf("SELECT name, description, price, amount FROM item_info ORDER BY price %s", sortForm)
//	exec, err := db.Query(command)
//
//	if err != nil {
//		log.Fatalf("Error ", err)
//	}
//	printItems(exec)
//}

func getPwd(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func printItems(rows *sql.Rows) error {
	var name, description string
	var price, amount float64
	i := 0
	for rows.Next() {
		err := rows.Scan(&name, &description, &price, &amount)
		if err != nil {
			return err
		}
		i++
		fmt.Printf("Name: %s, Description: %s, Price: %f, Amount: %f\n", name, description, price, amount)
	}
	if i == 0 {
		fmt.Println("No such items!")
	}
	return nil
}
