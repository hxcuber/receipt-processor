PROJECT_NAME := receipt-processor
COMPOSE := docker compose -f deploy/compose.yaml -p ${PROJECT_NAME}

local-build:
	go build ./cmd/app/main.go
local-run:
	go run ./cmd/app/main.go
local-test:
	go test ./...
clean:
	go clean && rm main

docker-build:
	docker build -f build/Dockerfile -t ${PROJECT_NAME}-api:latest .
docker-run:
	${COMPOSE} up app
docker-test:
	${COMPOSE} up test
docker-teardown:
	${COMPOSE} down

build-local: local-build
build-docker: docker-build
run-local: local-run
run-docker: docker-run
test-local: local-test
test-docker: docker-test