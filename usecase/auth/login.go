package auth

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"github.com/Fiddler25/ddd-sample-app/domain/repository/auth"
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

func (u LoginUsecase) Execute(req LoginRequest) model.User {
	user := auth.NewLoginRepository(u.db).GetByEmail(req.Email)

	return user
}
