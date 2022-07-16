package auth

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"github.com/Fiddler25/ddd-sample-app/domain/repository"
	"github.com/Fiddler25/ddd-sample-app/domain/vo"
	"github.com/Fiddler25/ddd-sample-app/sdk/cookie"
	"github.com/Fiddler25/ddd-sample-app/sdk/hash"
	"github.com/Fiddler25/ddd-sample-app/sdk/password"
	"github.com/Fiddler25/ddd-sample-app/sdk/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
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

	token := vo.NewRandomToken()
	user.SetRememberDigest(token)
	uRepo.UpdateRememberDigest(&user)

	cookie.Set(c, hash.Generate(strconv.Itoa(int(user.ID))), string(token))
	if err := bcrypt.CompareHashAndPassword([]byte(user.RememberDigest), []byte(token)); err != nil {
		return nil
	}

	return &user
}
