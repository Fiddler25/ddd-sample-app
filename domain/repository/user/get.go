package user

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"gorm.io/gorm"
	"log"
)

type GetRepository struct {
	db *gorm.DB
}

func NewGetRepository(db *gorm.DB) GetRepository {
	return GetRepository{db: db}
}

func (r GetRepository) ByUserID(userID int) model.User {
	var ret model.User
	if err := r.db.Find(&ret, userID).Error; err != nil {
		log.Fatal(err)
	}

	return ret
}
