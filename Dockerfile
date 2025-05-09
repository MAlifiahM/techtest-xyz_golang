# Build stage
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN cp .env.example .env

# Build with static linking
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o main ./cmd/app/main.go

# Use distroless as minimal base image
FROM gcr.io/distroless/static-debian11

WORKDIR /app

# Copy the binary
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["/app/main"]