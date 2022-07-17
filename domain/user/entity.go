package user

import (
	"github.com/Fiddler25/ddd-sample-app/domain/vo"
	"github.com/Fiddler25/ddd-sample-app/sdk/hash"
)

type UserID int

type User struct {
	ID             UserID
	Name           string
	Email          string
	Password       vo.Password
	RememberDigest vo.Token
}

func NewCreate(email string, password vo.Password) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}

func NewUpdate(userID UserID, name string, email string, password vo.Password) *User {
	return &User{
		ID:       userID,
		Name:     name,
		Email:    email,
		Password: password,
	}
}

func (u *User) SetRememberDigest(token vo.Token) {
	hashed := hash.Generate(string(token))
	u.RememberDigest = vo.Token(hashed)
}
