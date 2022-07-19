package auth

import (
	"github.com/Fiddler25/ddd-sample-app/domain/user"
	"github.com/Fiddler25/ddd-sample-app/domain/vo"
)

type Email string

type User struct {
	UserID   user.UserID
	Email    Email
	Password vo.Password
}
