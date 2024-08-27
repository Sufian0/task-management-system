package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/Sufian0/task-management-system/internal/database"
)

type Task struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func CreateTask(title, description string) (*Task, error) {
	task := &Task{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		Status:      "pending",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err := database.DB.Exec(
		"INSERT INTO tasks (id, title, description, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)",
		task.ID, task.Title, task.Description, task.Status, task.CreatedAt, task.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func GetAllTasks() ([]Task, error) {
	rows, err := database.DB.Query("SELECT id, title, description, status, created_at, updated_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func GetTaskByID(id uuid.UUID) (*Task, error) {
	var t Task
	err := database.DB.QueryRow("SELECT id, title, description, status, created_at, updated_at FROM tasks WHERE id = $1", id).
		Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func UpdateTask(id uuid.UUID, title, description, status string) (*Task, error) {
	task, err := GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	task.Title = title
	task.Description = description
	task.Status = status
	task.UpdatedAt = time.Now()

	_, err = database.DB.Exec(
		"UPDATE tasks SET title = $1, description = $2, status = $3, updated_at = $4 WHERE id = $5",
		task.Title, task.Description, task.Status, task.UpdatedAt, task.ID,
	)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func DeleteTask(id uuid.UUID) error {
	_, err := database.DB.Exec("DELETE FROM tasks WHERE id = $1", id)
	return err
}