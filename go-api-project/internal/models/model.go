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

// Follower represents a follower relationship between users.
type Follower struct {
    ID       uint `gorm:"primaryKey"`
    UserID   uint `json:"id_user"`
    FollowerID uint `json:"id_folower"`
}

// Like represents a like on a post or comment.
type Like struct {
    ID        uint `gorm:"primaryKey"`
    PostID    uint `json:"id_post"`
    UserID    uint `json:"id_user"`
    CommentID uint `json:"id_comment"`
}

// Member represents a member of a discussion.
type Member struct {
    ID           uint `gorm:"primaryKey"`
    UserID       uint `json:"id_user"`
    DiscussionID uint `json:"id_discussion"`
}

// Comment represents a comment made by a user.
type Comment struct {
    gorm.Model
    UserID  uint   `json:"id_user"`
    Content string `json:"content_comment"`
}

// Discussion represents a discussion between users.
type Discussion struct {
    gorm.Model
    Name string `json:"name_discussion"`
}

// Message represents a message sent in a discussion.
type Message struct {
    gorm.Model
    UserID       uint   `json:"id_user"`
    DiscussionID uint   `json:"id_discussion"`
    Content      string `json:"content_message"`
}