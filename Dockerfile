# First stage: Build the program
FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o main .

# Second stage: Dev stage with Air
FROM golang:1.22-alpine AS dev

WORKDIR /app
RUN go install github.com/air-verse/air@latest
COPY . .
CMD ["air"]

# Third stage: Prod stage
FROM alpine:latest AS prod

WORKDIR /app
COPY --from=builder /app/main /app/main
EXPOSE 8080
CMD ["./main"]