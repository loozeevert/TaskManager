package models

import "time"

type TaskStatus string

const (
	StatusTodo       TaskStatus = "todo"
	StatusInProgress TaskStatus = "in_progress"
	StatusDone       TaskStatus = "done"
)

type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	Priority    int        `json:"priority"` // 1-5, где 5 - наивысший
	DueDate     *time.Time `json:"due_date,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// DTO для создания задачи
type CreateTaskRequest struct {
	Title       string     `json:"title" binding:"required,min=1,max=255"`
	Description string     `json:"description" binding:"max=1000"`
	Status      TaskStatus `json:"status" binding:"omitempty,oneof=todo in_progress done"`
	Priority    int        `json:"priority" binding:"omitempty,min=1,max=5"`
	DueDate     *time.Time `json:"due_date"`
}

// DTO для обновления задачи
type UpdateTaskRequest struct {
	Title       string     `json:"title" binding:"omitempty,min=1,max=255"`
	Description *string    `json:"description" binding:"omitempty,max=1000"`
	Status      TaskStatus `json:"status" binding:"omitempty,oneof=todo in_progress done"`
	Priority    *int       `json:"priority" binding:"omitempty,min=1,max=5"`
	DueDate     *time.Time `json:"due_date"`
}