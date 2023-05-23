include .env

dep:
	go mod tidy
	go mod vendor

# Use this only for development
dev:
	go build -o bin/exchequer app/main.go
	./bin/exchequer

build:
	set GOOS=linux && set GOARCH=amd64 && go build -o bin/exchequer app/main.go

docker-build:
	docker-compose up --build

docker-up:
	docker-compose up

db-status:
	goose -dir ${GOOSE_MIGRATION_DIR} ${GOOSE_DRIVER} ${GOOSE_DBSTRING} status

migrate-up:
	goose -dir ${GOOSE_MIGRATION_DIR} ${GOOSE_DRIVER} ${GOOSE_DBSTRING} up
	
migrate-down:
	goose -dir ${GOOSE_MIGRATION_DIR} ${GOOSE_DRIVER} ${GOOSE_DBSTRING} down
	
migrate-reset:
	goose -dir ${GOOSE_MIGRATION_DIR} ${GOOSE_DRIVER} ${GOOSE_DBSTRING} reset
	
migrate-redo:
	goose -dir ${GOOSE_MIGRATION_DIR} ${GOOSE_DRIVER} ${GOOSE_DBSTRING} redo