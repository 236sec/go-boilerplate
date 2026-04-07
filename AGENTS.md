# Project Guidelines

## Code Style
- **Language**: Go 1.25.1. Strictly follow standard Go idioms and gofmt formatting.
- **Validation**: Performed via go-playground/validator/v10 applied at the REST handler level (e.g., [src/rest/handlers/validator.common.go](src/rest/handlers/validator.common.go)).
- **Database Schema vs Domain Entities**: Structural database models are in src/models/, separate from pure domain entities in src/domain/.

## Architecture
- **Clean Architecture**: Presentation (src/rest/), Application (src/usecases/), Domain (src/domain/), Infrastructure (src/repo/). 
- **Dependency Injection**: Explicitly managed in [src/di/](src/di/usecase.di.go) using sync.OnceValue to provide singletons. Do not use reflection libraries or init() functions.
- For a detailed breakdown of layers, see [README.md](README.md#project-architecture).

## Build and Test
- **Server**: make serve
- **Test**: make test
- **Lint**: make lint
- **Generate Mocks**: make gen-mock (mismatched mocks will cause tests to fail or act unpredictably).
- For the full list of commands, see [README.md](README.md#key-commands).

## Conventions
- **Mocks**: Mocks reside in src/repo/mocks/. Always regenerate mocks before testing if interfaces change.
- **Swagger**: The source lives in [docs/openapi.yaml](docs/openapi.yaml). Do not edit the compiled file manually. Run make swagger-generate.
- **Pre-commit**: The .githooks/pre-commit hook runs go mod tidy, generates mocks, lints, and runs tests. Ensure your work passes make test and make lint locally before committing.
- **Environment**: You must copy .env.example to .env to load environment variables locally (it is git-ignored).
