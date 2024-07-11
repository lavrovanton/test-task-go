FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@v3

RUN go build -o main ./cmd/api/main.go

EXPOSE 8080

CMD ["./main"]