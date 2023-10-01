ifeq ($(POSTGRES_SETUP_STRING),)
	POSTGRES_SETUP_STRING := user=test password=test dbname=postgres host=localhost port=5432 sslmode=disable
endif

INTERNAL_PKG_PATH=$(CURDIR)/internal/pkg
MIGRATION_FOLDER=$(CURDIR)/migrations

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: migration-up
migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_STRING)" up

.PHONY: migration-down
migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_STRING)" down

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
#	rm -rf ./internal/pb
#	mkdir -p ./internal/pb
	protoc ./api/*.proto/*.proto \
               --go_out=./internal/pb \
               --go-grpc_out=./internal/pb \
               --proto_path=.