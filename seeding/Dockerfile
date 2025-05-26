FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o seed ./seeding/main.go
RUN chmod +x /app/seed
USER root

ENTRYPOINT ["/app/seed"]