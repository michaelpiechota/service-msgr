package messenger

import "time"

// User data model
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Message data model
type Message struct {
	ID      string    `json:"id"`
	UserID  int       `json:"user_id"`
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
}
