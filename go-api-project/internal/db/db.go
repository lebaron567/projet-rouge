package db

import (
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "go-api-project/internal/models"
)

var DB *gorm.DB

func Init() {
    var err error
    DB, err = gorm.Open(sqlite.Open("message.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }

    // AutoMigrate will create the tables, missing foreign keys, constraints, columns and indexes.
    err = DB.AutoMigrate(
        &models.User{},
        &models.Post{},
        &models.Follower{},
        &models.Like{},
        &models.Member{},
        &models.Comment{},
        &models.Discussion{},
        &models.Message{},
    )
    if err != nil {
        log.Fatalf("Error migrating the database: %v", err)
    }
}