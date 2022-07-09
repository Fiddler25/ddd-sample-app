package repository

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"gorm.io/gorm"
	"log"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}

func (r UserRepository) GetByUserID(userID int) model.User {
	var ret model.User
	if err := r.db.Find(&ret, userID).Error; err != nil {
		log.Fatal(err)
	}
	return ret
}

func (r UserRepository) GetByEmail(email string) model.User {
	var ret model.User
	if err := r.db.Where("email = ?", email).Take(&ret).Error; err != nil {
		log.Fatal(err)
	}
	return ret
}

func (r UserRepository) Create(ret *model.User) *model.User {
	if err := r.db.Create(&ret).Error; err != nil {
		log.Fatal(err)
	}
	return ret
}
