# Makefile to generate code using oapi-codegen

OPENAPI_FILE=openapi.yaml
OUTPUT_DIR=generated

# Ensure oapi-codegen is installed
OAPI_CODEGEN=$(shell go env GOPATH)/bin/oapi-codegen

.PHONY: all clean install-oapi-codegen

# Install oapi-codegen if not present
install-oapi-codegen:
	@which oapi-codegen || (echo "Installing oapi-codegen..." && go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest)

# Generate all code
all: install-oapi-codegen generate-types generate-server generate-client

# Generate types
generate-types:
	$(OAPI_CODEGEN) -generate types -o $(OUTPUT_DIR)/types.gen.go $(OPENAPI_FILE)

# Generate server
generate-server:
	$(OAPI_CODEGEN) -generate gin -o $(OUTPUT_DIR)/server.gen.go $(OPENAPI_FILE)

# Generate client
generate-client:
	$(OAPI_CODEGEN) -generate client -o $(OUTPUT_DIR)/client.gen.go $(OPENAPI_FILE)

# Clean generated files
clean:
	rm -rf $(OUTPUT_DIR)/*.gen.go
