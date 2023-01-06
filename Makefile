## progog: Generate the protog
protog:
	@echo "===> Generating protog files"
	go mod download
	protoc \
		--go_out=./src/shared/contracts/protog \
		--go-grpc_out=./src/shared/contracts/protog \
		./src/shared/contracts/protodefs/*.proto

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose -f ./build/docker-compose.yml up -d
	@echo "Docker images started!"

## build: stops docker-compose (if running), builds all projects and starts docker compose
build: build_account_service build_account_service_api 
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

start_db:
	@echo "Stopping docker images (if running...)"
	docker-compose -f ./build/docker-compose.yml down
	@echo "===> Starting PostgresSQL"	
	docker-compose -f ./build/docker-compose.yml up -d

stop_db:
	@echo "===> Stopping PostgresSQL"
	docker-compose -f ./build/docker-compose.yml down

start_grpc: build_account_service
	@echo "===> Starting account service"
	./build/account_service/account_service

start_api: build_account_service_api
	@echo "===> Starting account service"
	./build/account_service_api/account_service_api

build_account_service:
	@echo "Building account service binary..."
	env GOOS=linux CGO_ENABLED=0 go build -o ./build/account_service  ./src/cmd/account_service
	@echo "Done!"

## build_account_api: builds the broker binary as a linux executable
build_account_service_api:
	@echo "Building account api binary..."	
	env GOOS=linux CGO_ENABLED=0 go build -o ./build/account_service_api  ./src/cmd/account_service_api
	@echo "Done!"	
	