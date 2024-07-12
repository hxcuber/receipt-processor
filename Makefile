COMPOSE := docker compose -f deploy/compose.yaml

local-build:
	go build ./cmd/app/main.go
local-run:
	go run ./cmd/app/main.go
local-test:
	go test ./...
clean:
	go clean && rm main

docker-build:
	docker build -f build/Dockerfile -t receipt-processor:latest .
docker-run:
	${COMPOSE} run --service-ports --rm -w /app app sh -c "go run cmd/app/main.go"
docker-test:
	${COMPOSE} -f deploy/compose.yaml run --service-ports --rm -w /app app sh -c "go test ./..."
docker-teardown:
	${COMPOSE} down

build-local: local-build
build-docker: docker-build
run-local: local-run
run-docker: docker-run
test-local: local-test
test-docker: docker-test