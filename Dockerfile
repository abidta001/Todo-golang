
FROM golang:1.23 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o todo-api

FROM alpine:latest
WORKDIR /app


RUN apk --no-cache add ca-certificates

COPY --from=builder /app/todo-api .

EXPOSE 3000

CMD ["./todo-api"]