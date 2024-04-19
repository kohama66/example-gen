package entity

type User struct {
	ID        int64  `json:"id"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}
