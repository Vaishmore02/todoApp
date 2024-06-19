package model

import "time"

var (
	ToDo      Status
	Inprogess Status
	Done      Status
)

type TodoTask struct {
	Id          string    `json:"id"`
	UserId      string    `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Status string
