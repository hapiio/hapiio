package db

import (
	"fmt"

	"github.com/hapiio/hapiio/core"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// connection as per config driver pending
// using sqllite for the development
func Connect(cnf core.Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(cnf.Database.Name), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Info: database connection established")
	return db
}

func Migrate(items interface{}, db *gorm.DB) bool {
	err := db.AutoMigrate(items)
	if err != nil {
		fmt.Print(err)
		return false
	}
	return true
}
