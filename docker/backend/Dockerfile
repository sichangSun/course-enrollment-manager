FROM golang:1.22.3-alpine3.19 AS builder

WORKDIR /app

COPY ./backend/ ./

RUN go mod tidy

RUN go build -o main main.go

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]