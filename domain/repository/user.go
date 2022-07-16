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

func (r UserRepository) GetByUserID(userID model.UserID) model.User {
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

func (r UserRepository) Update(ret *model.User) *model.User {
	if err := r.db.Updates(&ret).Error; err != nil {
		log.Fatal(err)
	}
	return ret
}

func (r UserRepository) Delete(userID model.UserID) {
	if err := r.db.Delete(&model.User{}, userID).Error; err != nil {
		log.Fatal(err)
	}
}

func (r UserRepository) UpdateRememberDigest(user *model.User) *model.User {
	if err := r.db.
		Model(&user).
		Where("id = ?", user.ID).
		Update("remember_digest", user.RememberDigest).Error; err != nil {
		log.Fatal(err)
	}
	return user
}
