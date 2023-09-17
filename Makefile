# 检查是否提供了 appName 参数，如果没有则使用默认值 go-gin-web
ifeq ($(appName),)
appName := go-gin-web
endif
# 检查是否提供了 version 参数，如果没有则使用默认值 1.0.0
ifeq ($(version),)
version := latest
endif


.PHONY: init
init:
	go mod tidy

# 本地打包
.PHONY: build
build:
	go build -ldflags="-s -w" -o ./bin/server ./cmd/server

# docker 打包镜像
.PHONY: docker-build
docker-build:
	docker build -f deploy/build/Dockerfile -t $(appName):$(version) .

# docker-compose 打包镜像
.PHONY: docker-compose-build
docker-compose-build:
	docker-compose -f deploy/build/docker-compose.yml build

# docker-compose 运行
.PHONY: docker-compose-up
docker-compose-up:
	docker-compose -f deploy/docker-compose/docker-compose.yml  -p go-gin-web up -d

# docker-compose 运行
.PHONY: docker-compose-down
docker-compose-down:
	docker-compose -f deploy/docker-compose/docker-compose.yml down



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