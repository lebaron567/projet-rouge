package database

import (
	"api/database/dbmodel"
	"log"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	// Effectuer les migrations pour toutes les tables
	err := db.AutoMigrate(
		&dbmodel.User{},
		&dbmodel.Follower{},
		&dbmodel.Post{},
		&dbmodel.Like{},
		&dbmodel.Comment{},
	)

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	} else {
		log.Println("Migration successful")
	}
}
