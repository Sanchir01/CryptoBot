PHONY:
SILENT:
include .env

build:
	go build -o ./.bin/main ./cmd/main/main.go

run: @export $(shell sed 's/=.*//' .env) && \
 	@echo "TOKEN_BOT: $(TOKEN_BOT)" && \
    @echo "API_KEY_BINANCE: $(API_KEY_BINANCE)" && \
    @echo "SECRET_KEY_BINANCE: $(SECRET_KEY_BINANCE)"\
	build && \
	./.bin/main

build-image:
	docker build -t cryptobot-dockerfile .
start-container:
	docker run --name cryptobot-test -p 80:80 --env-file .env cryptobot-dockerfile