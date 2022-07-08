package user

import (
	"github.com/Fiddler25/ddd-sample-app/domain/user"
	"gorm.io/gorm"
	"log"
)

func AddPassword(db *gorm.DB) {
	m := db.Migrator()

	if ok := m.HasColumn(&user.User{}, "Password"); !ok {
		if err := m.AddColumn(&user.User{}, "Password"); err != nil {
			log.Fatal(err)
		}
	}
}
