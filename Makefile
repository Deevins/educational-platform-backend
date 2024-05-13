ifeq ($(POSTGRES_SETUP_STRING),)
	POSTGRES_SETUP_STRING := user=Shili password=postgres dbname=pg host=localhost port=5432 sslmode=disable
endif

INTERNAL_PKG_PATH=$(CURDIR)/internal/pkg
MIGRATION_FOLDER=$(CURDIR)/migrations
EDUCATIONAL_PLATFORM_MAIN = $(CURDIR)/cmd/educational-platform/main.go

.PHONY: run
run:
	go build "$(EDUCATIONAL_PLATFORM_MAIN)"/cmd/educational-platform

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: migration-up
migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_STRING)" up

.PHONY: migration-down
migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_STRING)" down

.PHONY: migration-reset
migration-reset:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_STRING)" reset

.PHONY: compose-up
compose-up:
	docker-compose build

.PHONY: compose-rm
compose-rm:
	docker-compose down

.PHONY: unit-test
unit-tests:
	go test .\internal\app\handlers -v

.PHONY: integration-test
unit-tests:
	go test .\test\ -v

.PHONY: proto
proto:
	protoc -I ./api/ -I./api/google/api \
        --go_out=./internal/pb --go_opt=paths=source_relative \
        --go-grpc_out=./internal/pb --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=./internal/pb --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out=./internal/pb --openapiv2_opt=logtostderr=true \
        ./api/users/users.proto
