all: build run

build:
	templ generate &\
	go build -o bin/app

run:
	./bin/app
