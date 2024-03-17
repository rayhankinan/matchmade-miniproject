include .env
export

PWD = $(shell pwd)

run:
	go run main.go start --port=8080

migrate-up:
	go run main.go migrate-up

migrate-down:
	go run main.go migrate-down