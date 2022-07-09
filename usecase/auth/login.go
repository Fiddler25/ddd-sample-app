package auth

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"github.com/Fiddler25/ddd-sample-app/domain/repository"
	"github.com/Fiddler25/ddd-sample-app/sdk"
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
	user := repository.NewUser(u.db).GetByEmail(req.Email)
	if !sdk.IsSamePassword(string(user.Password), req.Password) {
		return nil
	}
	sdk.Login(c, user.ID)

	return &user
}
