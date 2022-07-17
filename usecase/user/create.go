package user

import (
	"fmt"
	"github.com/Fiddler25/ddd-sample-app/domain/user"
	"github.com/Fiddler25/ddd-sample-app/domain/vo"
	"gorm.io/gorm"
)

type CreateRequest struct {
	Email                string `json:"email" validate:"required,max=255,email"`
	Password             string `json:"password" validate:"required,min=6"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,min=6"`
}

type CreateUsecase struct {
	db *gorm.DB
}

func NewCreateUsecase(db *gorm.DB) CreateUsecase {
	return CreateUsecase{db: db}
}

func (uc CreateUsecase) Execute(req CreateRequest) (*user.User, error) {
	hash := vo.NewPassword(req.Password)
	u := user.NewCreate(req.Email, hash)

	uRepo := user.NewRepository(uc.db)
	if usr, _ := uRepo.GetByEmail(u.Email); usr != nil {
		return nil, fmt.Errorf("メールアドレスが重複しています：%s", u.Email)
	}
	uRepo.Create(u)

	return u, nil
}
