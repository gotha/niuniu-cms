package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(c *Config) (*gorm.DB, error) {
	dsn := c.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}

	migrateDB := os.Getenv("MIGRATE_DB")
	if migrateDB == "true" {
		q := `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`
		res := db.Exec(q)
		if res.Error != nil {
			return nil, fmt.Errorf("could not create uuid-ossp extension")
		}

		err = db.AutoMigrate(&Document{}, &Tag{})
		if err != nil {
			return nil, fmt.Errorf("could not migrate database schema: %w", err)
		}
	}

	return db, nil
}
