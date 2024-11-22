all: build run

build:
	@templ generate
	@go build ./cmd/api/main.go

run:
	@./main
