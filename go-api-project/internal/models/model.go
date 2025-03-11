package models

// User represents a user in the system.
type User struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// Post represents a post created by a user.
type Post struct {
    ID      string `json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
    UserID  string `json:"user_id"`
}