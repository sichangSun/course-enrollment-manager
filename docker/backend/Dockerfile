FROM golang:1.17.0-alpine3.13 AS builder

WORKDIR /app

COPY ./backend/main.go ./

RUN go build -o main main.go

FROM alpine:3.13

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]