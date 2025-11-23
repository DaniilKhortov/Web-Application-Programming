# --- Етап 1: збірка ---
FROM golang:1.21-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /queue-app ./cmd/web

# --- Етап 2: фінальний мінімальний образ ---
FROM scratch
COPY --from=builder /queue-app /queue-app
EXPOSE 8080
ENTRYPOINT ["/queue-app"]
