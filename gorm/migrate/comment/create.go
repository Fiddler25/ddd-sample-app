package comment

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"gorm.io/gorm"
	"log"
)

func Create(db *gorm.DB) {
	m := db.Migrator()

	if ok := m.HasTable(&model.Comment{}); !ok {
		if err := m.CreateTable(&model.Comment{}); err != nil {
			log.Fatal(err)
		}
	}
}
