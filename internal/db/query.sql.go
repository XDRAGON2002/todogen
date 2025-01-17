// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package db

import (
	"context"
)

const addTodo = `-- name: AddTodo :one
INSERT INTO todos (
    id, title, isCompleted
) VALUES (
    NULL, ?, ?
)
RETURNING id, title, iscompleted
`

type AddTodoParams struct {
	Title       string
	Iscompleted int64
}

func (q *Queries) AddTodo(ctx context.Context, arg AddTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, addTodo, arg.Title, arg.Iscompleted)
	var i Todo
	err := row.Scan(&i.ID, &i.Title, &i.Iscompleted)
	return i, err
}

const deleteTodoById = `-- name: DeleteTodoById :exec
DELETE FROM todos
WHERE id = ?
`

func (q *Queries) DeleteTodoById(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTodoById, id)
	return err
}

const getTodoById = `-- name: GetTodoById :one
SELECT id, title, iscompleted FROM todos
WHERE id = ?
`

func (q *Queries) GetTodoById(ctx context.Context, id int64) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getTodoById, id)
	var i Todo
	err := row.Scan(&i.ID, &i.Title, &i.Iscompleted)
	return i, err
}

const getTodos = `-- name: GetTodos :many
SELECT id, title, iscompleted FROM todos
`

func (q *Queries) GetTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, getTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(&i.ID, &i.Title, &i.Iscompleted); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTodoById = `-- name: UpdateTodoById :one
UPDATE todos
SET title = ?, isCompleted = ?
WHERE id = ?
RETURNING id, title, iscompleted
`

type UpdateTodoByIdParams struct {
	Title       string
	Iscompleted int64
	ID          int64
}

func (q *Queries) UpdateTodoById(ctx context.Context, arg UpdateTodoByIdParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, updateTodoById, arg.Title, arg.Iscompleted, arg.ID)
	var i Todo
	err := row.Scan(&i.ID, &i.Title, &i.Iscompleted)
	return i, err
}
