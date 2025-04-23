package model

import "gorm.io/gorm"

// Post represents a post made by a user
// @Description Post represents a post made by a user
type Post struct {
	gorm.Model
	ID       uint      `json:"id"`
	IDUser   int       `json:"id_user"`
	Title    string    `json:"title_post"`
	Content  string    `json:"description_post"`
	IsStory  bool      `json:"isstory_post"`
	User     User      `json:"user"`
	Likes    []Like    `json:"likes"`
	Comments []Comment `json:"comments"`
}
