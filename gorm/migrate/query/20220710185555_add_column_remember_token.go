package query

import (
	"gorm.io/gorm"
)

func AddColumnRememberToken(db *gorm.DB) {
	var count int
	db.Raw("SELECT count(*) FROM INFORMATION_SCHEMA.columns WHERE TABLE_NAME = 'users' AND COLUMN_NAME = 'remember_token'").Row().Scan(&count)

	if count > 0 {
		return
	}
	db.Exec("ALTER TABLE users ADD COLUMN remember_token varchar(255) NOT NULL")
}
