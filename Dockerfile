# First stage: Build the program
FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o main .

# Second stage: Run the program
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main /app/main

EXPOSE 8080
CMD ["./main"]