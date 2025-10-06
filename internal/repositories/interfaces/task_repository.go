package interfaces

import (
	"context"
	"task-manager/internal/models"
)

type TaskRepository interface {
	Create(ctx context.Context, task *models.Task) error
	GetByID(ctx context.Context, id int) (*models.Task, error)
	GetAll(ctx context.Context) ([]*models.Task, error)
	Update(ctx context.Context, id int, task *models.Task) error
	Delete(ctx context.Context, id int) error
	GetByCompleted(ctx context.Context, completed bool) ([]*models.Task, error)
}
