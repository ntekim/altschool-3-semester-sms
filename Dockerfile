#build stage
FROM golang:1.20-alpine3.17 AS builder

WORKDIR /app

COPY . .

RUN go build -o main main.go

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -command="./app"

FROM alpine:latest

WORKDIR /app

COPY  --from=builder /app/main .

CMD ["/app/main"]