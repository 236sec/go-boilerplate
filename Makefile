# Path variable
MIGRATIONS_PATH = ./src/migrations

serve:
	go run ./src/main.go

migration-generate:
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(name)

migration-up:
	go run ./src/cmd/migrate.go up

migration-down:
	go run ./src/cmd/migrate.go down
