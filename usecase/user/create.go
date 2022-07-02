package user

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"github.com/Fiddler25/ddd-sample-app/domain/repository/user"
	"github.com/Fiddler25/ddd-sample-app/sdk"
	"gorm.io/gorm"
)

type CreateUsecase struct {
	db *gorm.DB
}

func NewCreateUsecase(db *gorm.DB) CreateUsecase {
	return CreateUsecase{db: db}
}

func (u CreateUsecase) Execute(data model.User) model.User {
	rawPassword := data.Password
	hash := sdk.HashPassword(rawPassword)
	data.Password = hash

	return user.NewCreateRepository(u.db).Create(data)
}
