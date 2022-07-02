package user

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"gorm.io/gorm"
	"log"
)

type CreateRepository struct {
	db *gorm.DB
}

func NewCreateRepository(db *gorm.DB) CreateRepository {
	return CreateRepository{db: db}
}

func (r CreateRepository) Create(data model.User) model.User {
	if err := r.db.Create(&data).Error; err != nil {
		log.Fatal(err)
	}

	return data
}
