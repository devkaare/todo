all: build test

build:
	@echo "Building..."
	@templ generate
	@go build -o ./main ./cmd/api/main.go

run:
	@echo "Running..."
	@go run ./cmd/api/main.go

docker-run:
	@if docker compose up --build 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up --build; \
	fi

docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

test:
	@echo "Testing..."
	@go test ./... -v

clean: 
	@echo "Cleaning..."
	@go mod tidy
	@rm ./main
	@rm -rf ./views/*_templ.go

watch:
	@echo "Watching..."
	@air

.PHONY: all build run test clean watch docker-run docker-down
