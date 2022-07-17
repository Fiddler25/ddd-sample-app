package user

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"github.com/Fiddler25/ddd-sample-app/domain/repository"
	"github.com/Fiddler25/ddd-sample-app/domain/vo"
	"gorm.io/gorm"
)

type CreateRequest struct {
	Email                string `json:"email" validate:"required,max=255,email"`
	Password             string `json:"password" validate:"required,min=6"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,min=6"`
}

type CreateUsecase struct {
	db *gorm.DB
}

func NewCreateUsecase(db *gorm.DB) CreateUsecase {
	return CreateUsecase{db: db}
}

func (u CreateUsecase) Execute(req CreateRequest) *model.User {
	hash := vo.NewPassword(req.Password)
	user := model.NewCreateUser(req.Email, hash)

	repository.NewUser(u.db).Create(user)

	return user
}
