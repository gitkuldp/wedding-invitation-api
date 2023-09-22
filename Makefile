serve:
	go run cmd/main.go

migrate:
	go run cmd/migrate/migrate.go

.PHONY:docs
docs:
	@swag fmt
	@swag init -g cmd/main.go