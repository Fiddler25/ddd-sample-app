package user

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (r Repository) ByUserID(userID int) model.User {
	var ret model.User
	if err := r.db.Find(&ret, userID).Error; err != nil {
		log.Fatal(err)
	}

	return ret
}
