package auth

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"github.com/Fiddler25/ddd-sample-app/domain/repository"
	"github.com/Fiddler25/ddd-sample-app/sdk/cookie"
	"github.com/Fiddler25/ddd-sample-app/sdk/password"
	"github.com/Fiddler25/ddd-sample-app/sdk/session"
	"github.com/labstack/echo/v4"
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

func (u LoginUsecase) Execute(c echo.Context, req LoginRequest) *model.User {
	uRepo := repository.NewUser(u.db)
	user := uRepo.GetByEmail(req.Email)
	if !password.IsSame(string(user.Password), req.Password) {
		return nil
	}
	session.Login(c, user.ID)

	user.SetRememberToken()
	uRepo.UpdateRememberToken(&user)

	cookie.Set(c, string(user.RememberToken))

	return &user
}
