FROM golang:1.27-alpine as builder

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN go build -o bin/todo-list-api ./cmd/todo-list-api

FROM alpine

COPY --from=builder /app/bin/todo-list-api /todo-list-api
COPY --from=builder /app/.env .env

CMD [ "/todo-list-api" ]