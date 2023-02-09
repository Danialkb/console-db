package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Authorization interface {
	SignIn(string, string) error
}

type AuthorizationService struct{}

func (s *AuthorizationService) SignIn(login, password string) error {
	db, _ := initDB()
	command := fmt.Sprintf("SELECT password FROM user_info WHERE login='%s';", login)
	exec, err := db.Query(command)

	if err != nil {
		return err
	}

	if exec.Next() {
		var passwordDB string
		exec.Scan(&passwordDB)
		if err := ComparePasswords(passwordDB, []byte(password)); err != nil {
			return err
		} else {
			return nil
		}
	}
	db.Close()
	return nil
}

func ComparePasswords(hashedPwd string, plainPwd []byte) error {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return err
	}
	return nil
}
