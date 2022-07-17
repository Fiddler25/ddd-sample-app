package auth

import (
	"fmt"
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

func (uc LoginUsecase) Execute(req LoginRequest) (*user.User, error) {
	repo := user.NewRepository(uc.db)

	if u, err := repo.GetByEmail(req.Email); err != nil {
		return nil, fmt.Errorf("メールアドレスが存在しません")
	} else if !password.IsSame(string(u.Password), req.Password) {
		return nil, fmt.Errorf("パスワードが一致しません")
	} else {
		return u, nil
	}
}
