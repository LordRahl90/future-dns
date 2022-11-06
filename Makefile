.PHONY: bi

all: bi
	@docker-compose up

test:
	@go test ./... --cover

start:
	@go run ./cmd/

build-image:
	@docker build -t gcr.io/neurons-be-test/dns:latest .


bi:build-image