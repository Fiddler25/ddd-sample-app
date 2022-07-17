package user

import (
	"github.com/Fiddler25/ddd-sample-app/domain/user"
	"github.com/Fiddler25/ddd-sample-app/domain/vo"
	"gorm.io/gorm"
)

type UpdateRequest struct {
	ID                   user.UserID `json:"id" validate:"required"`
	Name                 string      `json:"name" validate:"max=50"`
	Email                string      `json:"email" validate:"omitempty,max=255,email"`
	Password             string      `json:"password" validate:"omitempty,min=6"`
	PasswordConfirmation string      `json:"password_confirmation" validate:"omitempty,min=6"`
}

type UpdateUsecase struct {
	db *gorm.DB
}

func NewUpdateUsecase(db *gorm.DB) UpdateUsecase {
	return UpdateUsecase{db: db}
}

func (uc UpdateUsecase) Execute(req UpdateRequest) *user.User {
	hash := vo.NewPassword(req.Password)
	u := user.NewUpdate(req.ID, req.Name, req.Email, hash)

	user.NewRepository(uc.db).Update(u)

	return u
}
