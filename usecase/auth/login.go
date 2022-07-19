package auth

import (
	"errors"
	"github.com/Fiddler25/ddd-sample-app/domain/auth"
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

func (uc LoginUsecase) Execute(req LoginRequest) (*auth.User, error) {
	repo := auth.NewRepository(uc.db)

	if usr, err := repo.GetByEmail(auth.Email(req.Email)); err != nil {
		return nil, errors.New("メールアドレスが存在しません")
	} else if !auth.IsSamePassword(string(usr.Password), req.Password) {
		return nil, errors.New("パスワードが一致しません")
	} else {
		return usr, nil
	}
}
