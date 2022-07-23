package query

import (
	"gorm.io/gorm"
)

func AlterColumnName(db *gorm.DB) {
	db.Exec("ALTER TABLE users ALTER name SET DEFAULT ''")
}
