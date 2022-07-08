package user

import (
	"github.com/Fiddler25/ddd-sample-app/domain/user"
	"gorm.io/gorm"
	"log"
)

func Create(db *gorm.DB) {
	m := db.Migrator()

	if ok := m.HasTable(&user.User{}); !ok {
		if err := m.CreateTable(&user.User{}); err != nil {
			log.Fatal(err)
		}
	}
}
