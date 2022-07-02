package user

import (
	"github.com/Fiddler25/ddd-sample-app/domain/user"
	"gorm.io/gorm"
)

type Usecase struct {
	db *gorm.DB
}

func NewGet(db *gorm.DB) Usecase {
	return Usecase{db: db}
}

func (u Usecase) Execute(userID int) user.User {
	return user.NewUserRepository(u.db).ByUserID(userID)
}
