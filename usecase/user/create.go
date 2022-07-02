package user

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"github.com/Fiddler25/ddd-sample-app/domain/repository/user"
	"github.com/Fiddler25/ddd-sample-app/sdk"
	"gorm.io/gorm"
)

type CreateRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type CreateUsecase struct {
	db *gorm.DB
}

func NewCreateUsecase(db *gorm.DB) CreateUsecase {
	return CreateUsecase{db: db}
}

func (u CreateUsecase) Execute(req CreateRequest) model.User {
	hash := sdk.HashPassword(req.Password)
	data := model.NewUser(req.Email, hash)

	return user.NewCreateRepository(u.db).Create(data)
}
