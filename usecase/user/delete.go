package user

import (
	"github.com/Fiddler25/sample-app/domain/user"
	"gorm.io/gorm"
)

type DeleteUsecase struct {
	db *gorm.DB
}

func NewDeleteUsecase(db *gorm.DB) DeleteUsecase {
	return DeleteUsecase{db: db}
}

func (uc DeleteUsecase) Execute(userID user.UserID) {
	user.NewRepository(uc.db).Delete(userID)
}
