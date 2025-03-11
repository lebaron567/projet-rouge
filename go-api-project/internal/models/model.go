package models

import "gorm.io/gorm"

// User represents a user in the system.
type User struct {
    gorm.Model
    LastName       string `json:"lastename_user"`
    FirstName      string `json:"firstname_user"`
    Email          string `json:"email_user"`
    Pseudo         string `json:"pseudo_user"`
    Birthdate      string `json:"birthdate"`
    Password       string `json:"password_user"`
    IsPrivate      bool   `json:"isprivate_user"`
    ProfilePicture string `json:"profilpicture_user"`
    WantsNotify    bool   `json:"wantsnotify_user"`
}

// Post represents a post created by a user.
type Post struct {
    gorm.Model
    UserID    uint   `json:"user_id"`
    Title     string `json:"title_post"`
    Content   string `json:"content"`
    IsStory   bool   `json:"isstory_post"`
}