# 1claw-go-sdk Makefile
# Regenerates the internal OpenAPI client from the spec.

SPEC_PATH ?= ../1claw-openapi-spec/openapi.yaml
OUTPUT_PATH = ./internal/openapi

.PHONY: generate
generate:
	@if [ ! -f "$(SPEC_PATH)" ]; then \
		echo "Error: OpenAPI spec not found at $(SPEC_PATH)"; \
		echo "Set SPEC_PATH if the spec is elsewhere, e.g. make generate SPEC_PATH=/path/to/openapi.yaml"; \
		exit 1; \
	fi
	npx @openapitools/openapi-generator-cli generate \
		-i "$(SPEC_PATH)" \
		-g go \
		-o "$(OUTPUT_PATH)" \
		--additional-properties=packageName=openapi,outputAsLibrary=true
	@# Remove generated go.mod so internal/openapi is part of parent module
	@rm -f $(OUTPUT_PATH)/go.mod $(OUTPUT_PATH)/go.sum
	@echo "Generated client at $(OUTPUT_PATH)"

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	go build ./...
