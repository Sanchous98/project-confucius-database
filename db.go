package database

import (
	"github.com/Sanchous98/project-confucius-base/stdlib"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// Database represents GORM-to-Confucius bridge
type Database struct {
	gorm.DB
	Log stdlib.Logger `inject:""`
}

func (d *Database) Constructor() {
	db, err := gorm.Open(mysql.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	d.Log.Alert(err)
	d.DB = *db
}
