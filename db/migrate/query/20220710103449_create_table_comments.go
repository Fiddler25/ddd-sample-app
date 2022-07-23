package query

import (
	"gorm.io/gorm"
)

func CreateTableComments(db *gorm.DB) {
	db.Exec(`CREATE TABLE IF NOT EXISTS comments (
		id bigint AUTO_INCREMENT,
		body longtext NOT NULL,
		user_id bigint NOT NULL,
		created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (id),
		INDEX idx_user_id_created_at (user_id, created_at),
		CONSTRAINT fk_comments_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	)`)
}
