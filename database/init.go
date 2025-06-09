package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init(dsn string) (*gorm.DB, error) {
	dialector := sqlite.Open(dsn)

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("unable to open database, %w", err)
	}

	err = db.AutoMigrate(&Subnet{}, &IP{})
	if err != nil {
		return nil, fmt.Errorf("unable to migrate database models, %w", err)
	}

	return db, nil
}
