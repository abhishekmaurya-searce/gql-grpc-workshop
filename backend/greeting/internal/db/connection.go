package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewConnection(dsn string, schema interface{}) (DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return DB{}, fmt.Errorf("error in connecting to database: %s", err)
	}
	err = db.AutoMigrate(schema)
	if err != nil {
		return DB{}, fmt.Errorf("error in migrating database schema: %s", err)
	}
	return DB{DB: db}, nil
}
