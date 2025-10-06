package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"task-manager/internal/database"
	"task-manager/internal/models"
)

type TaskRepository struct {
	db *database.DB
}

func NewTaskRepository(db *database.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(ctx context.Context, task *models.Task) error {
	query := `
		INSERT INTO tasks (title, description, completed, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)
	`

	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now

	result, err := r.db.ExecContext(ctx, query,
		task.Title,
		task.Description,
		task.Completed,
		task.CreatedAt,
		task.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("erro ao criar task: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("erro ao obter ID da task criada: %v", err)
	}

	task.ID = int(id)
	return nil
}

func (r *TaskRepository) GetByID(ctx context.Context, id int) (*models.Task, error) {
	query := `
		SELECT id, title, description, completed, created_at, updated_at
		FROM tasks 
		WHERE id = ?
	`

	task := &models.Task{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Completed,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("task com ID %d não encontrada", id)
		}
		return nil, fmt.Errorf("erro ao buscar task: %v", err)
	}

	return task, nil
}

func (r *TaskRepository) GetAll(ctx context.Context) ([]*models.Task, error) {
	query := `
		SELECT id, title, description, completed, created_at, updated_at
		FROM tasks 
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar tasks: %v", err)
	}
	defer rows.Close()

	var tasks []*models.Task
	for rows.Next() {
		task := &models.Task{}
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear task: %v", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *TaskRepository) Update(ctx context.Context, id int, task *models.Task) error {
	query := `
		UPDATE tasks 
		SET title = ?, description = ?, completed = ?, updated_at = ?
		WHERE id = ?
	`

	task.UpdatedAt = time.Now()

	result, err := r.db.ExecContext(ctx, query,
		task.Title,
		task.Description,
		task.Completed,
		task.UpdatedAt,
		id,
	)
	if err != nil {
		return fmt.Errorf("erro ao atualizar task: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar linhas afetadas: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("task com ID %d não encontrada", id)
	}

	return nil
}

func (r *TaskRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM tasks WHERE id = ?`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("erro ao deletar task: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar linhas afetadas: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("task com ID %d não encontrada", id)
	}

	return nil
}

func (r *TaskRepository) GetByCompleted(ctx context.Context, completed bool) ([]*models.Task, error) {
	query := `
		SELECT id, title, description, completed, created_at, updated_at
		FROM tasks 
		WHERE completed = ?
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, completed)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar tasks por status: %v", err)
	}
	defer rows.Close()

	var tasks []*models.Task
	for rows.Next() {
		task := &models.Task{}
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear task: %v", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
