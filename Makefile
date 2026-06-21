-include ./.env

MIGRATION_PATH=db/migration

migrate-create:
	@migrate create -ext sql -dir $(MIGRATION_PATH) -seq create_$(NAME)_table

migrate-up:
	@migrate -database $(DB_URL) -path $(MIGRATION_PATH) up

migrate-down:
	@migrate -database $(DB_URL) -path $(MIGRATION_PATH) down

migrate-force:
	@migrate -database $(DB_URL) -path $(MIGRATION_PATH) force $(VERSION)

sqlc:
	sqlc generate

test:
	go test -v ./...

.PHONY: migrate-create migrate-up migrate-down migrate-force sqlc test