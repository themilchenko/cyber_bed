package models

type AuthUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       uint64 `json:"userID"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

type Username struct {
	ID       uint64
	Username string
}

type UsersInfo struct {
	UserID   uint64
	Password string
	Avatar   string
}

type UserID struct {
	ID uint64 `json:"userID"`
}
