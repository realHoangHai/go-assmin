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
	swag init -d ./cmd,./internal/service,./internal/model,./internal/common -g main.go --output docs/qltb

build:
	GOOS=linux GOARCH=amd64 go build -o bin/assmin ./cmd/