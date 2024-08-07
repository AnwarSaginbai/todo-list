package ports

import (
	"context"
	"github.com/AnwarSaginbai/todo-list/internal/application/domain"
)

type APIPort interface {
	SaveTaskToDB(ctx context.Context, todo *domain.Todo) (*domain.Todo, error)
	GetAllTaskFromDB(ctx context.Context) ([]*domain.Todo, error)
	DeleteTaskFromDB(ctx context.Context, id int) error
	UpdateTaskToDB(ctx context.Context, id int, todo domain.Todo) error
}
