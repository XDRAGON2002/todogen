package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"example.com/todogo/internal/api"
	"example.com/todogo/internal/db"
)

type todosService struct {
	queries *db.Queries
}

func (t *todosService) GetTodos(ctx context.Context) ([]api.Todo, error) {
	items, err := t.queries.GetTodos(context.Background())
	if err != nil {
		return nil, err
	}

	var todos []api.Todo
	for _, item := range items {
		todos = append(todos, api.Todo{
			ID:          api.NewOptInt64(item.ID),
			Title:       item.Title,
			IsCompleted: api.NewOptBool(bool(item.Iscompleted != 0)),
		})
	}

	return todos, nil
}

func (t *todosService) GetTodoById(ctx context.Context, params api.GetTodoByIdParams) (api.GetTodoByIdRes, error) {
	item, err := t.queries.GetTodoById(context.Background(), params.TodoId)
	if err != nil {
		return nil, err
	}

	return &api.Todo{
		ID:          api.NewOptInt64(item.ID),
		Title:       item.Title,
		IsCompleted: api.NewOptBool(bool(item.Iscompleted != 0)),
	}, nil
}

func (t *todosService) AddTodo(ctx context.Context, req *api.Todo) (*api.Todo, error) {
	isCompleted := 0
	if req.GetIsCompleted().Value {
		isCompleted = 1
	}

	item, err := t.queries.AddTodo(context.Background(), db.AddTodoParams{
		Title:       req.Title,
		Iscompleted: int64(isCompleted),
	})
	if err != nil {
		return nil, err
	}

	return &api.Todo{
		ID:          api.NewOptInt64(item.ID),
		Title:       item.Title,
		IsCompleted: api.NewOptBool(bool(item.Iscompleted != 0)),
	}, nil
}

func (t *todosService) UpdateTodoById(ctx context.Context, req *api.Todo, params api.UpdateTodoByIdParams) (*api.Todo, error) {
	isCompleted := 0
	if req.GetIsCompleted().Value {
		isCompleted = 1
	}

	item, err := t.queries.UpdateTodoById(context.Background(), db.UpdateTodoByIdParams{
		Title:       req.Title,
		Iscompleted: int64(isCompleted),
		ID:          params.TodoId,
	})
	if err != nil {
		return nil, err
	}

	return &api.Todo{
		ID:          api.NewOptInt64(item.ID),
		Title:       item.Title,
		IsCompleted: api.NewOptBool(bool(item.Iscompleted != 0)),
	}, nil
}

func (t *todosService) DeleteTodoById(ctx context.Context, params api.DeleteTodoByIdParams) error {
	err := t.queries.DeleteTodoById(context.Background(), params.TodoId)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	dbCtx, err := sql.Open("sqlite3", "todos.db")
	if err != nil {
		log.Fatal(err)
	}

	queries := db.New(dbCtx)

	service := &todosService{
		queries: queries,
	}

	srv, err := api.NewServer(service)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening on port 8080...")
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
