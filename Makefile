docker.run:
	docker run -d --name my-postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres-secret -e POSTGRES_DB=assmin postgres:15-alpine

run:
	go run ./cmd

wire.install:
	go install github.com/google/wire/cmd/wire@latest

wire.gen:
	wire ./cmd/...

ent.install:
	go install entgo.io/ent/cmd/ent@latest

ent.init:
	go run -mod=mod entgo.io/ent/cmd/ent init --target internal/ent/schema User


ent.gen:
	go generate ./internal/ent/...

swag.install:
	go get github.com/swaggo/swag/cmd/swag
	go install github.com/swaggo/swag/cmd/swag
	go get github.com/swaggo/gin-swagger
	go get github.com/swaggo/files

swag.gen:
	swag init -d ./cmd,./internal/service,./internal/model,./internal/common/response,./internal/common/errors -g main.go --output docs/swagger

build:
	GOOS=linux GOARCH=amd64 go build -o bin/assmin ./cmd/