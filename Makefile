.PHONY:

build:
	go build -o ./.bin/bot app/cmd/bot/main.go

run: build
	./.bin/bot