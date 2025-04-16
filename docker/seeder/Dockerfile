FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY internal/seeder/ ./internal/seeder/
COPY cmd/seeder/main.go ./cmd/seeder/main.go

RUN go build -o seeder ./cmd/seeder
RUN chmod +x /app/seeder

CMD ["./seeder"]
