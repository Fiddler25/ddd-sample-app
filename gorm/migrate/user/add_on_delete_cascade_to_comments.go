package user

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"gorm.io/gorm"
	"log"
)

func AddOnDeleteCascadeToUser(db *gorm.DB) {
	m := db.Migrator()

	if ok := m.HasConstraint(&model.User{}, "Comments"); !ok {
		if err := m.CreateConstraint(&model.User{}, "Comments"); err != nil {
			log.Fatal(err)
		}
	}
}
