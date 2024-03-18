FROM golang:1.22.1 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o todo-api main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/todo-api .

EXPOSE 8080

CMD ["./todo-api"]
