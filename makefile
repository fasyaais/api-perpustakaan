dev:
	go run ./cmd/app/main.go

migrate:
	migrate create -ext sql -dir internal/migrations -seq $(name)