FROM golang:1.21-alpine3.17 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main server/main.go

FROM alpine:3.17

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main","server"]
