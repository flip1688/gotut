package service

import (
	"errors"
	"gotut/database"
)

// Service type
type Service struct {
}

var (
	ErrInvalidLogin = errors.New("invalid username or password")
)

// CreateUser to create a new user
func (*Service) CreateUser(u *database.User) error {
	db, err := database.Connect()
	if err != nil {
		return errors.New("Can not create db connection")
	}
	defer db.Close()
	if !db.HasTable(&database.User{}) {
		db = db.CreateTable(&database.User{})
	}
	db.Create(u)
	return nil
}

// VerifyUserPassword for login function
func (*Service) VerifyUserPassword(username string, password string) (*database.User, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, errors.New("Can not create db connection")
	}
	defer db.Close()
	findUser := database.User{
		Username: username,
		Password: password,
	}
	var foundUser database.User
	db.Where(&findUser).First(&foundUser)
	if foundUser.ID == 0 {
		return nil, ErrInvalidLogin
	}
	return &foundUser, nil
}
