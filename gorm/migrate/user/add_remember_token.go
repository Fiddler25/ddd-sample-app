package user

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"gorm.io/gorm"
	"log"
)

func AddRememberToken(db *gorm.DB) {
	m := db.Migrator()

	if ok := m.HasColumn(&model.User{}, "RememberToken"); !ok {
		if err := m.AddColumn(&model.User{}, "RememberToken"); err != nil {
			log.Fatal(err)
		}
	}
}
