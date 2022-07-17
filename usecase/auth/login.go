package auth

import (
	"github.com/Fiddler25/ddd-sample-app/domain/user"
	"github.com/Fiddler25/ddd-sample-app/sdk/password"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUsecase struct {
	db *gorm.DB
}

func NewLoginUsecase(db *gorm.DB) LoginUsecase {
	return LoginUsecase{db: db}
}

func (uc LoginUsecase) Execute(req LoginRequest) *user.User {
	repo := user.NewRepository(uc.db)
	u := repo.GetByEmail(req.Email)
	if !password.IsSame(string(u.Password), req.Password) {
		return nil
	}

	return &u
}
