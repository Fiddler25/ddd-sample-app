package auth

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"github.com/Fiddler25/ddd-sample-app/domain/repository"
	"github.com/Fiddler25/ddd-sample-app/domain/vo"
	"github.com/Fiddler25/ddd-sample-app/sdk/cookie"
	"github.com/Fiddler25/ddd-sample-app/sdk/hash"
	"github.com/Fiddler25/ddd-sample-app/sdk/password"
	"github.com/Fiddler25/ddd-sample-app/sdk/session"
	"github.com/Fiddler25/ddd-sample-app/sdk/session/entity"
	sModel "github.com/Fiddler25/ddd-sample-app/sdk/session/model"
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

func (uc LoginUsecase) Execute(c echo.Context, req LoginRequest) *model.User {
	repo := repository.NewUser(uc.db)
	u := repo.GetByEmail(req.Email)
	if !password.IsSame(string(u.Password), req.Password) {
		return nil
	}
	session.Login(c, u.ID)
	sessID := entity.NewSessionID()
	sess := sModel.NewSession(sessID, u.ID)

	token := vo.NewRandomToken()
	u = updateRememberDigest(u, repo, token)

	cookie.Set(c, hash.Generate(strconv.Itoa(int(u.ID))), string(token))
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
