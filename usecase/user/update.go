package user

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"github.com/Fiddler25/ddd-sample-app/domain/vo"
	"gorm.io/gorm"
)

type UpdateRequest struct {
	ID                   int    `json:"id" validate:"required"`
	Name                 string `json:"name" validate:"max=50"`
	Email                string `json:"email" validate:"omitempty,max=255,email"`
	Password             string `json:"password" validate:"omitempty,min=6"`
	PasswordConfirmation string `json:"password_confirmation" validate:"omitempty,min=6"`
}

type UpdateUsecase struct {
	db *gorm.DB
}

func NewUpdateUsecase(db *gorm.DB) UpdateUsecase {
	return UpdateUsecase{db: db}
}

func (u UpdateUsecase) Execute(req UpdateRequest) *model.User {
	hash := vo.NewPassword(req.Password)
	user := model.NewUpdateUser(req.ID, req.Name, req.Email, hash)

	return user
}
