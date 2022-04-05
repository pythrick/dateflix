LINUX_AMD64 = GOOS=linux GOARCH=amd64 CGO_ENABLED=1 GO111MODULE=on

build:
	@$(LINUX_AMD64) go build -a -v -tags musl -o dateflix-api main.go

start:
	@docker-compose --env-file .local.env -f docker-compose.yml up -d

stop:
	@docker-compose --env-file .local.env -f docker-compose.yml down
