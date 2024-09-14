run:
	@go run main.go

build:
	@go build -o bin/server

ogen:
	@go run github.com/ogen-go/ogen/cmd/ogen@latest --target internal/api --clean openapi.yaml

sqlc:
	@go run github.com/sqlc-dev/sqlc/cmd/sqlc@latest generate

atlas:
	@atlas schema apply --url "sqlite://todos.db" --dev-url "sqlite://todos.db" --to "file://schema.sql" --auto-approve
