package user

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"github.com/Fiddler25/ddd-sample-app/domain/repository"
	"gorm.io/gorm"
)

type GetUsecase struct {
	db *gorm.DB
}

func NewGetUsecase(db *gorm.DB) GetUsecase {
	return GetUsecase{db: db}
}

func (u GetUsecase) Execute(userID model.UserID) model.User {
	return repository.NewUser(u.db).GetByUserID(userID)
}
