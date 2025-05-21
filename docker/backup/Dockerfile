FROM golang:1.24-alpine

RUN apk add --no-cache postgresql15-client gzip bash tzdata

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o backup-runner ./cmd/backup

COPY internal/backup/backup.sh ./backup.sh
RUN chmod +x ./backup.sh ./backup-runner
CMD ["./backup-runner"]
