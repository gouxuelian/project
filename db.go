package main

import (
	"github.com/jinzhu/gorm"
	// nolint: gotype
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB is the DB that will performs all operation
type DB struct {
	DB *gorm.DB
}

// NewDB returns a new DB connection
func newDB(path string) (*DB, error) {
	// connect to the example db, create if it doesn't exist
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Employee{})

	return &DB{db}, nil
}
