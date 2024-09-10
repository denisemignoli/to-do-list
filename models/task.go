package models

type Task struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
	UserID    int64  `json:"user_id"`
}
