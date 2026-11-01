package db

import (
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewDb() (*gorm.DB){
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail to connect to database")
	}
	return db
}