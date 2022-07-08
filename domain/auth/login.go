package auth

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"gorm.io/gorm"
	"log"
)

type LoginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(db *gorm.DB) LoginRepository {
	return LoginRepository{db: db}
}

func (r LoginRepository) GetByEmail(email string) model.User {
	var ret model.User
	if err := r.db.Where("email = ?", email).Take(&ret).Error; err != nil {
		log.Fatal(err)
	}

	return ret
}
