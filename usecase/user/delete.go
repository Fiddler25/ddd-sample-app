package user

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"github.com/Fiddler25/ddd-sample-app/domain/repository"
	"gorm.io/gorm"
)

type DeleteUsecase struct {
	db *gorm.DB
}

func NewDeleteUsecase(db *gorm.DB) DeleteUsecase {
	return DeleteUsecase{db: db}
}

func (u DeleteUsecase) Execute(userID model.UserID) {
	repository.NewUser(u.db).Delete(userID)
}