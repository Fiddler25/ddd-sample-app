package auth

import "github.com/Fiddler25/ddd-sample-app/domain/user"

type Email string

type User struct {
	ID       user.UserID
	Email    Email
	Password Password
}

func New(email Email, password Password) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}
