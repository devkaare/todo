FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/a-h/templ/cmd/templ@latest && \
	templ generate
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/api/main.go

EXPOSE ${PORT}
CMD ["./main"]

