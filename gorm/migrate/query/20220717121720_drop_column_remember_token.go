package query

import (
	"gorm.io/gorm"
)

func DropColumnRememberDigest(db *gorm.DB) {
	var count int
	db.Raw("SELECT count(*) FROM INFORMATION_SCHEMA.columns WHERE TABLE_NAME = 'users' AND COLUMN_NAME = 'remember_digest'").Row().Scan(&count)

	if count == 0 {
		return
	}
	db.Exec("ALTER TABLE users DROP COLUMN remember_digest")
}
