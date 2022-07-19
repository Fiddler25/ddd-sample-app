package user

import (
	"github.com/Fiddler25/ddd-sample-app/domain/user"
	"gorm.io/gorm"
)

type GetUsecase struct {
	db *gorm.DB
}

func NewGetUsecase(db *gorm.DB) GetUsecase {
	return GetUsecase{db: db}
}

func (uc GetUsecase) Execute(userID user.UserID) (user.User, error) {
	return user.NewRepository(uc.db).GetByUserID(userID)
}
