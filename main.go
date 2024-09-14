package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	api "example.com/todogo/internal/api"
)

type todosService struct {
	todos []api.Todo
	id    int64
	mux   sync.Mutex
}

func (t *todosService) GetTodos(ctx context.Context) ([]api.Todo, error) {
	t.mux.Lock()
	defer t.mux.Unlock()

	return t.todos, nil
}

func (t *todosService) GetTodoById(ctx context.Context, params api.GetTodoByIdParams) (api.GetTodoByIdRes, error) {
	t.mux.Lock()
	defer t.mux.Unlock()

	for _, todo := range t.todos {
		if val, ok := todo.ID.Get(); ok && val == int(params.TodoId) {
			return &todo, nil
		}
	}
	return &api.GetTodoByIdNotFound{}, nil
}

func (t *todosService) AddTodo(ctx context.Context, req *api.Todo) (*api.Todo, error) {
	t.mux.Lock()
	defer t.mux.Unlock()

	todo := req
	todo.ID.SetTo(int(t.id))
	t.id++
	t.todos = append(t.todos, *todo)

	return todo, nil
}

func (t *todosService) DeleteTodoById(ctx context.Context, params api.DeleteTodoByIdParams) error {
	t.mux.Lock()
	defer t.mux.Unlock()

	idxToDelete := -1
	for idx, todo := range t.todos {
		if val, ok := todo.ID.Get(); ok && val == int(params.TodoId) {
			idxToDelete = idx
			break
		}
	}

	if idxToDelete != -1 {
		t.todos = append(t.todos[:idxToDelete], t.todos[idxToDelete+1:]...)
	}

	return nil
}

func main() {
	service := &todosService{}

	srv, err := api.NewServer(service)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening on port 8080...")
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
