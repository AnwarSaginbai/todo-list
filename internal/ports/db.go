package ports

import (
	"context"
	"github.com/AnwarSaginbai/todo-list/internal/application/domain"
)

type DBPort interface {
	InsertTask(ctx context.Context, todo *domain.Todo) error
	GetAllTask(ctx context.Context) ([]*domain.Todo, error)
	DeleteTaskByID(ctx context.Context, id int) error
	UpdateTask(ctx context.Context, id int, todo domain.Todo) error
}
