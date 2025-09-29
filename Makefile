# Path variable
MIGRATIONS_PATH = ./src/migrations
OPENAPI_PATH = ./docs/openapi.yaml
SWAGGER_COMPILED_PATH = ./docs/compile/swagger.yaml

serve:
	go run ./src/main.go

migration-generate:
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(name)

migration-up:
	go run ./src/cmd/migrate.go up

migration-down:
	go run ./src/cmd/migrate.go down

install-swagger-generate:
	npm install -g swagger-cli

swagger-generate:
	swagger-cli bundle $(OPENAPI_PATH) --outfile $(SWAGGER_COMPILED_PATH) --type yaml
