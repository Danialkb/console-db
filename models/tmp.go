package models

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func initDB() (*sql.DB, error) {
	db, err := sql.Open("postgres",
		"user=postgres password=Dankb2131193* host=localhost port=5432 dbname=assignment1_go sslmode=disable")

	if err != nil {
		return nil, err
	}

	return db, nil
}

//func printItems(exec *sql.Rows) error {
//
//	for exec.Next() {
//		var (
//			name        string
//			description string
//			price       float64
//			amount      int
//		)
//
//		err := exec.Scan(&name, &description, &price, &amount)
//
//		if err != nil {
//			log.Fatalf("ERROR ", err)
//			return err
//		}
//		fmt.Println(name, description, price, amount)
//	}
//	exec.Close()
//	return nil
//}

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

func print_all(exec *sql.Rows) {
	for exec.Next() {
		var (
			id          int
			name        string
			description string
			price       float64
			amount      int
			user_id     int
		)

		exec.Scan(&id, &name, &description, &price, &amount, &user_id)

		fmt.Println(name, description, price, amount)
	}
}
