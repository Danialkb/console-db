package models

import (
	"fmt"
	"time"
)

type Registration interface {
	Register(*User) bool
}

type RegistrationService struct{}

func (r *RegistrationService) Register(u *User) error {
	db, _ := initDB()
	command := fmt.Sprintf("SELECT * FROM user_info WHERE login='%s';", u.Login)
	exec, err := db.Query(command)

	if err != nil {
		return err
	}
	if exec.Next() {
		fmt.Println("User with such login already exists")
		return nil
	}

	if u.Password, err = getPwd(u.Password); err != nil {
		return err
	}

	if _, err := db.Exec("INSERT INTO user_info(login, password, created_at)VALUES ($1, $2, $3)",
		u.Login, u.Password, time.Now().UTC()); err != nil {
		return err
	}
	db.Close()
	return nil
}
