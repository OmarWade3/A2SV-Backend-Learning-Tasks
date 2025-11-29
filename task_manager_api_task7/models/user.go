package models

type User struct {
	ID       int    `json:"id"`
	Username string `josn:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
