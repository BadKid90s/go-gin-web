# 检查是否提供了 appName 参数，如果没有则使用默认值 go-gin-web
ifeq ($(appName),)
appName := go-gin-web
endif
# 检查是否提供了 version 参数，如果没有则使用默认值 1.0.0
ifeq ($(version),)
version := 1.0.0
endif


.PHONY: init
init:
	go mod tidy

.PHONY: build
build:
	go build -ldflags="-s -w" -o ./bin/server ./cmd/server

.PHONY: docker
docker:
	docker build -f deploy/build/Dockerfile -t $(appName):$(version) .

#.PHONY: bootstrap
#bootstrap:
#	cd ./deploy/docker-compose && docker compose up -d && cd ../../
#	go run ./cmd/migration
#	nunu run ./cmd/server
#
#.PHONY: mock
#mock:
#	mockgen -source=internal/service/user.go -destination test/mocks/service/user.go
#	mockgen -source=internal/repository/user.go -destination test/mocks/repository/user.go
#
#.PHONY: test
#test:
#	go test -coverpkg=./internal/handler,./internal/service,./internal/repository -coverprofile=./coverage.out ./test/server/...
#	go tool cover -html=./coverage.out -o coverage.html
#
#
#.PHONY: swag
#swag:
#	swag init  -g cmd/server/main.go -o ./docs --parseDependency