package auth

type Email string

type User struct {
	Email    Email
	Password Password
}

func New(email Email, password Password) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}
