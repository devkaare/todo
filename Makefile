all: build run

build:
	@echo "Building..."
	@templ generate
	@go build ./cmd/api/main.go

run:
	@echo "Running..."
	@./main

clean: 
	@echo "Cleaning..."
	@go mod tidy
	@rm ./main
	@rm -rf ./views/*_templ.go
