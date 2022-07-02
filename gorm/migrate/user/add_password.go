package user

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"gorm.io/gorm"
	"log"
)

func AddPassword(db *gorm.DB) {
	m := db.Migrator()

	if ok := m.HasColumn(&model.User{}, "Password"); !ok {
		if err := m.AddColumn(&model.User{}, "Password"); err != nil {
			log.Fatal(err)
		}
	}
}
