package user

import (
	"github.com/Fiddler25/ddd-sample-app/domain/user"
	"gorm.io/gorm"
)

type GetUsecase struct {
	db *gorm.DB
}

func NewGetUsecase(db *gorm.DB) GetUsecase {
	return GetUsecase{db: db}
}

func (u GetUsecase) Execute(userID int) user.User {
	return user.NewGetRepository(u.db).ByUserID(userID)
}
