include .env
export $(shell sed 's/=.*//' .env)

gen_docs:
	swag fmt && swag init -g cmd/api/main.go

run: gen_docs
	go run cmd/api/main.go
