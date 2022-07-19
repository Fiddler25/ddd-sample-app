package auth

import (
	"fmt"
	"github.com/Fiddler25/ddd-sample-app/domain/auth"
	"gorm.io/gorm"
)

type SignUpRequest struct {
	Email                string `json:"email" validate:"required,max=255,email"`
	Password             string `json:"password" validate:"required,min=6"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,min=6"`
}

type SignUpUsecase struct {
	db *gorm.DB
}

func NewSignUpUsecase(db *gorm.DB) SignUpUsecase {
	return SignUpUsecase{db: db}
}

func (uc SignUpUsecase) Execute(req SignUpRequest) (*auth.User, error) {
	hash := auth.NewPassword(req.Password)
	email := auth.Email(req.Email)
	usr := auth.New(email, hash)

	aRepo := auth.NewRepository(uc.db)
	if registeredUser, _ := aRepo.GetByEmail(usr.Email); registeredUser != nil {
		return nil, fmt.Errorf("メールアドレスが重複しています：%s", usr.Email)
	}
	aRepo.Create(usr)

	return usr, nil
}
