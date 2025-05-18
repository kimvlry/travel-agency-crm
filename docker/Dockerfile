FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o seeder ./cmd/seeder
RUN go build -o monitoring ./cmd/monitoring

RUN chmod +x /app/seeder /app/monitoring