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

func (uc LoginUsecase) Execute(c echo.Context, req LoginRequest) *model.User {
	repo := repository.NewUser(uc.db)
	u := repo.GetByEmail(req.Email)
	if !password.IsSame(string(u.Password), req.Password) {
		return nil
	}

	var oldSessionID session.ID
	if ck, err := c.Cookie(cookie.CookieNameSession); err != nil {
		oldSessionID = ""
	} else {
		oldSessionID = session.ID(ck.Value)
	}

	sRepo := session.NewRepository(uc.db)
	if oldSessionID != "" {
		if oldSession, err := sRepo.Find(oldSessionID); err != nil {
			return nil
		} else {
			sRepo.Delete(oldSession)
		}
	}

	sessionID := session.NewID()
	sess := session.NewModel(sessionID, u.ID)
	sRepo.Create(sess)

	cookie.Set(c, string(sess.SessionID))

	return &u
}
