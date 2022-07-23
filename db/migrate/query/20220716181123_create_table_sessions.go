package query

import (
	"gorm.io/gorm"
)

func CreateTableSessions(db *gorm.DB) {
	db.Exec(`CREATE TABLE IF NOT EXISTS sessions (
		session_id varchar(45) NOT NULL,
		user_id int(10) unsigned DEFAULT NULL,
		created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (session_id)
	)`)
}
