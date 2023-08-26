package dbutil

import (
	"context"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// TODO
func GormDB(ctx context.Context) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(os.Getenv("DB_CONNECTION_STRING")), &gorm.Config{})
	return db, err
}
