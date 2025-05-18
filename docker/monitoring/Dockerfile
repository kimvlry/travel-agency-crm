FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o metrics-exporter ./cmd/monitoring
RUN chmod +x /app/metrics-exporter

CMD [ "./metrics-exporter" ]
