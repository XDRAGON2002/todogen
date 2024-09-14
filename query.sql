-- name: GetTodos :many
SELECT * FROM todos;

-- name: GetTodoById :one
SELECT * FROM todos
WHERE id = ?;

-- name: AddTodo :one
INSERT INTO todos (
    id, title, isCompleted
) VALUES (
    NULL, ?, ?
)
RETURNING *; 

-- name: UpdateTodoById :one
UPDATE todos
SET title = ?, isCompleted = ?
WHERE id = ?
RETURNING *;

-- name: DeleteTodoById :exec
DELETE FROM todos
WHERE id = ?;
