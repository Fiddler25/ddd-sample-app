package user

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"github.com/Fiddler25/ddd-sample-app/domain/repository/user"
	"gorm.io/gorm"
)

type GetUsecase struct {
	db *gorm.DB
}

func NewGetUsecase(db *gorm.DB) GetUsecase {
	return GetUsecase{db: db}
}

func (u GetUsecase) Execute(userID int) model.User {
	return user.NewGetRepository(u.db).ByUserID(userID)
}
