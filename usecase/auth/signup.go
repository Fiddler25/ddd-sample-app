package auth

import "gorm.io/gorm"

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

func (uc SignUpUsecase) Execute(req SignUpRequest) {

}