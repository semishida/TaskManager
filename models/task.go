package models

type Task struct {
	ID          int    `json:"id"`
	UserID      uint   `json:"user_id"`
	Group       string `json:"group"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
	Status      string `json:"status"`
}

type Report struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
