.PHONY: bi

all: bi
	@docker-compose up

test:
	@go test ./... --cover

start:
	@go run ./cmd/

build-image:
	@docker build -t gcr.io/neurons-be-test/dns:latest .

generate:
	protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false \
	--go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/request.proto

bi:build-image