package db

import (
	"context"
	"fmt"
	"github.com/AnwarSaginbai/todo-list/internal/application/domain"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

type Adapter struct {
	db *sqlx.DB
}

func NewAdapter(dsn string) (*Adapter, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err = db.PingContext(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return &Adapter{db: db}, nil
}
func (a *Adapter) InsertTask(ctx context.Context, todo *domain.Todo) error {
	query := `
        INSERT INTO todo (task, done, created_at, updated_at, deleted_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id;
    `

	var id int
	err := a.db.QueryRowContext(ctx, query, todo.Task, todo.Done, todo.CreatedAt, todo.UpdatedAt, todo.DeletedAt).Scan(&id)
	if err != nil {
		return err
	}

	todo.ID = id // Обновляем ID в переданном объекте
	return nil
}

func (a *Adapter) GetAllTask(ctx context.Context) ([]*domain.Todo, error) {
	query := `SELECT * FROM todo;`

	rows, err := a.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*domain.Todo

	for rows.Next() {
		var todo domain.Todo
		if err := rows.Scan(&todo.ID, &todo.Task, &todo.Done, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (a *Adapter) DeleteTaskByID(ctx context.Context, id int) error {
	query := `DELETE FROM todo WHERE id = $1`

	result, err := a.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no task found with id %d", id)
	}

	return nil
}

func (a *Adapter) UpdateTask(ctx context.Context, id int, todo domain.Todo) error {
	query := `
        UPDATE todo
        SET task = $1, done = $2, updated_at = $3
        WHERE id = $4;
    `

	result, err := a.db.ExecContext(ctx, query, todo.Task, todo.Done, todo.UpdatedAt, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no task found with id %d", id)
	}

	return nil
}
