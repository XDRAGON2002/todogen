// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// AddTodo implements addTodo operation.
	//
	// Add a new todo to the list.
	//
	// POST /todos
	AddTodo(ctx context.Context, req *Todo) (*Todo, error)
	// DeleteTodoById implements deleteTodoById operation.
	//
	// Deletes a single todo.
	//
	// DELETE /todos/{todoId}
	DeleteTodoById(ctx context.Context, params DeleteTodoByIdParams) error
	// GetTodoById implements getTodoById operation.
	//
	// Returns a single todo.
	//
	// GET /todos/{todoId}
	GetTodoById(ctx context.Context, params GetTodoByIdParams) (GetTodoByIdRes, error)
	// GetTodos implements getTodos operation.
	//
	// Returns all todos.
	//
	// GET /todos
	GetTodos(ctx context.Context) ([]Todo, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
