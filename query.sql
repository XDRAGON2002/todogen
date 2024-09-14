-- name: GetTodos :many
SELECT * FROM todos;

-- name: GetTodoById :one
SELECT * FROM todos
WHERE id = ?;

-- name: AddTodo :one
INSERT INTO todos (
    id
) VALUES (
    NULL
)
RETURNING *; 

-- name: DeleteTodoById :exec
DELETE FROM todos
WHERE id = ?;
