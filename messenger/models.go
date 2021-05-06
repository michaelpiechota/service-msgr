package messenger

// User data model
type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Message data model
type Message struct {
	ID      string `json:"id"`
	UserID  int64  `json:"user_id"`
	Message string `json:"message"`
}
