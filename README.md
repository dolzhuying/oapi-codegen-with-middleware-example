# Sample API with Gin, OpenAPI withMiddleware

This project demonstrates how to use `oapi-codegen` with the Gin web framework and custom middleware in Golang.

## API Endpoints

- `GET /items`: Retrieve a list of items
- `POST /items`: Create a new item
- `GET /items/{id}`: Retrieve a specific item by its ID

## Middleware

- **Global middleware**: Logs before and after each request.
- **Before and after request middleware**: Applied to route groups.
- **Admin-specific middleware**: Applied to `POST /items` routes.

## Project Structure

- `main.go`: Main application file.
- `openapi.yaml`: OpenAPI specification.
- `generated/`: Directory for generated server code from `oapi-codegen`.

## Makefile

The Makefile provides tasks to generate code using `oapi-codegen`:

- `make install-oapi-codegen`: Install `oapi-codegen` if not already installed.
- `make generate-types`: Generate type definitions.
- `make generate-server`: Generate server code.
- `make generate-client`: Generate client code.
- `make clean`: Remove generated files.
