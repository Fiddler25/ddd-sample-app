package user

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"github.com/Fiddler25/ddd-sample-app/domain/repository/user"
	"gorm.io/gorm"
)

type Usecase struct {
	db *gorm.DB
}

func GetUsecase(db *gorm.DB) Usecase {
	return Usecase{db: db}
}

func (u Usecase) Execute(userID int) model.User {
	return user.GetRepository(u.db).ByUserID(userID)
}
