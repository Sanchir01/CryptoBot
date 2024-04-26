FROM golang:1.22.2-alpine3.18 AS build

COPY . /github.com/Sanchir01/CryptoBot
WORKDIR  /github.com/Sanchir01/CryptoBot

RUN go mod download

RUN go build -o ./bin/main ./cmd/main/main.go


FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/Sanchir01/CryptoBot/bin/main .
COPY --from=0 /github.com/Sanchir01/CryptoBot/configs configs/


EXPOSE 80

CMD ["./main"]