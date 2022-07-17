package auth

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"github.com/Fiddler25/ddd-sample-app/domain/repository"
	"github.com/Fiddler25/ddd-sample-app/domain/vo"
	"github.com/Fiddler25/ddd-sample-app/sdk/cookie"
	"github.com/Fiddler25/ddd-sample-app/sdk/password"
	"github.com/Fiddler25/ddd-sample-app/sdk/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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
	session.Login(c, u.ID)
	var oldSessionID session.ID
	if ck, err := c.Cookie(cookie.CookieNameSession); err != nil {
		oldSessionID = ""
	} else {
		oldSessionID = session.ID(ck.Value)
	}
	sessID := session.NewID()
	sess := session.NewModel(sessID, u.ID)
	session.NewRepository(uc.db).Create(sess)

	cookie.Set(c, string(sess.SessionID))

	token := vo.NewRandomToken()
	u = updateRememberDigest(u, repo, token)

	if !isAuthenticated(u.RememberDigest, token) {
		return nil
	}

	return &u
}

func updateRememberDigest(user model.User, repo repository.UserRepository, token vo.Token) model.User {
	user.SetRememberDigest(token)
	repo.UpdateRememberDigest(&user)

	return user
}

func isAuthenticated(digest, token vo.Token) bool {
	err := bcrypt.CompareHashAndPassword([]byte(digest), []byte(token))
	return err == nil
}
