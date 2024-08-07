package api

import (
	"context"
	"fmt"
	"github.com/AnwarSaginbai/todo-list/internal/application/domain"
	"github.com/AnwarSaginbai/todo-list/internal/ports"
)

type Service struct {
	db ports.DBPort
}

func NewService(db ports.DBPort) *Service {
	return &Service{db: db}
}

func (s *Service) SaveTaskToDB(ctx context.Context, todo *domain.Todo) (*domain.Todo, error) {
	err := s.db.InsertTask(ctx, todo)
	if err != nil {
		return &domain.Todo{}, fmt.Errorf("failed to save in db: %v", err)
	}
	return todo, nil
}

func (s *Service) GetAllTaskFromDB(ctx context.Context) ([]*domain.Todo, error) {
	data, err := s.db.GetAllTask(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all tasks from db: %v\n", err)
	}
	return data, nil
}

func (s *Service) DeleteTaskFromDB(ctx context.Context, id int) error {
	if err := s.db.DeleteTaskByID(ctx, id); err != nil {
		return fmt.Errorf("failed to delete the task: %v\n", err)
	}
	return nil
}

func (s *Service) UpdateTaskToDB(ctx context.Context, id int, todo domain.Todo) error {
	if err := s.db.UpdateTask(ctx, id, todo); err != nil {
		return fmt.Errorf("failed to update the task: %v\n", err)
	}
	return nil
}
