package repository

import (
	"database/sql"
	"fmt"
	"project/internal/models"
)

type TaskRepository interface {
	GetAll() ([]models.Task, error)
	GetByID(id int) (*models.Task, error)
	Create(task *models.Task) error
	Update(id int, task *models.Task) error
	Delete(id int) error
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{db: db}
}

// GetAll возвращает все задачи
func (r *taskRepository) GetAll() ([]models.Task, error) {
	query := `SELECT id, title, description, status, priority, due_date, created_at, updated_at 
	          FROM tasks ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %w", err)
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&task.ID, &task.Title, &task.Description, &task.Status,
			&task.Priority, &task.DueDate, &task.CreatedAt, &task.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return tasks, nil
}

// GetByID возвращает задачу по ID
func (r *taskRepository) GetByID(id int) (*models.Task, error) {
	query := `SELECT id, title, description, status, priority, due_date, created_at, updated_at 
	          FROM tasks WHERE id = $1`

	row := r.db.QueryRow(query, id)

	var task models.Task
	err := row.Scan(
		&task.ID, &task.Title, &task.Description, &task.Status,
		&task.Priority, &task.DueDate, &task.CreatedAt, &task.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("task with id %d not found", id)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %w", err)
	}

	return &task, nil
}

// Create создает новую задачу
func (r *taskRepository) Create(task *models.Task) error {
	query := `INSERT INTO tasks (title, description, status, priority, due_date) 
	          VALUES ($1, $2, $3, $4, $5) 
	          RETURNING id, created_at, updated_at`

	err := r.db.QueryRow(
		query,
		task.Title,
		task.Description,
		task.Status,
		task.Priority,
		task.DueDate,
	).Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}

	return nil
}

// Update обновляет существующую задачу
func (r *taskRepository) Update(id int, task *models.Task) error {
	query := `UPDATE tasks 
	          SET title = $1, description = $2, status = $3, priority = $4, due_date = $5 
	          WHERE id = $6 
	          RETURNING updated_at`

	err := r.db.QueryRow(
		query,
		task.Title,
		task.Description,
		task.Status,
		task.Priority,
		task.DueDate,
		id,
	).Scan(&task.UpdatedAt)

	if err == sql.ErrNoRows {
		return fmt.Errorf("task with id %d not found", id)
	}
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	task.ID = id
	return nil
}

// Delete удаляет задачу по ID
func (r *taskRepository) Delete(id int) error {
	query := `DELETE FROM tasks WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("task with id %d not found", id)
	}

	return nil
}
