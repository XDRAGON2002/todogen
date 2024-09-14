// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// AddTodo implements addTodo operation.
//
// Add a new todo to the list.
//
// POST /todos
func (UnimplementedHandler) AddTodo(ctx context.Context, req *Todo) (r *Todo, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteTodoById implements deleteTodoById operation.
//
// Deletes a single todo.
//
// DELETE /todos/{todoId}
func (UnimplementedHandler) DeleteTodoById(ctx context.Context, params DeleteTodoByIdParams) error {
	return ht.ErrNotImplemented
}

// GetTodoById implements getTodoById operation.
//
// Returns a single todo.
//
// GET /todos/{todoId}
func (UnimplementedHandler) GetTodoById(ctx context.Context, params GetTodoByIdParams) (r GetTodoByIdRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetTodos implements getTodos operation.
//
// Returns all todos.
//
// GET /todos
func (UnimplementedHandler) GetTodos(ctx context.Context) (r []Todo, _ error) {
	return r, ht.ErrNotImplemented
}