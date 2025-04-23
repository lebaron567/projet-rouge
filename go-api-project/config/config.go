package config

import (
	"log"

	"api/database/dbmodel"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	DB                 *gorm.DB
	UserRepository     dbmodel.UserRepository
	PostRepository     dbmodel.PostRepository
	CommentRepository  dbmodel.CommentRepository
	LikeRepository     dbmodel.LikeRepository
	FollowerRepository dbmodel.FollowerRepository
}

func New() (*Config, error) {
	db, err := gorm.Open(sqlite.Open("yline.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Impossible de se connecter à la base de données : %v", err)
		return nil, err
	}
	if err := db.AutoMigrate(&dbmodel.User{}, &dbmodel.Post{}, &dbmodel.Comment{}, &dbmodel.Like{}, &dbmodel.Follower{}); err != nil {
		log.Fatalf("Erreur lors de la migration : %v", err)
		return nil, err
	}

	userRepo := dbmodel.NewUserRepository(db)
	postRepo := dbmodel.NewPostRepository(db)
	commentRepo := dbmodel.NewCommentRepository(db)
	likeRepo := dbmodel.NewLikeRepository(db)
	followRepo := dbmodel.NewFollowerRepository(db)

	config := Config{
		DB:                 db,
		UserRepository:     userRepo,
		PostRepository:     postRepo,
		CommentRepository:  commentRepo,
		LikeRepository:     likeRepo,
		FollowerRepository: followRepo,
	}

	return &config, nil
}
