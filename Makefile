VERSION ?= $(shell git describe --match 'v[0-9]*' --tags --always)
BUILD_FOLDER=./build/dateflix-api


version:
	@echo $(VERSION)

format:
	gofmt -w -l .

lint:
	@$(MAKE) format
	@./scripts/lint_go_mods.sh

test:
	@go test -v ./...

download:
	@echo Download go.mod dependencies
	@go mod download

tools: download
	@echo Installing tools from tools.go
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

docker-build:
	@docker-compose build

docker-run:
	@docker-compose up 

install:
	@go mod tidy

run:
	@go run cmd/dateflix-api/main.go
