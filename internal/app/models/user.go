package models

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UsersBooks struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	BookID string `json:"book_id"`
}
