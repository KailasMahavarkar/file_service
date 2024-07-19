.PHONY: build start

build:
	go build -o ./bin/app main.go

start:
	nodemon --signal SIGTERM
