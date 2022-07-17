package user

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (r Repository) GetByUserID(userID UserID) (User, error) {
	var ret User
	if err := r.db.Where("id = ?", userID).Take(&ret).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, err
		}
	}
	return ret, nil
}

func (r Repository) GetByEmail(email string) User {
	var ret User
	if err := r.db.Where("email = ?", email).Take(&ret).Error; err != nil {
		log.Fatal(err)
	}
	return ret
}

func (r Repository) Create(ret *User) *User {
	if err := r.db.Create(&ret).Error; err != nil {
		log.Fatal(err)
	}
	return ret
}

func (r Repository) Update(ret *User) *User {
	if err := r.db.Updates(&ret).Error; err != nil {
		log.Fatal(err)
	}
	return ret
}

func (r Repository) Delete(userID UserID) {
	if err := r.db.Delete(&User{}, userID).Error; err != nil {
		log.Fatal(err)
	}
}
