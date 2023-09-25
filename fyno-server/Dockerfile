# Build stage
FROM golang:1.20 AS builder

ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./bin/app ./cmd

# Final stage
FROM alpine:3.14

WORKDIR /

COPY --from=builder /app/bin/app .
COPY db/migrations /db/migrations

CMD ["./app"]
