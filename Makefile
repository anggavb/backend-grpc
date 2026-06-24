-include ./app.env

MIGRATION_PATH=db/migration

migrate-create:
	@migrate create -ext sql -dir $(MIGRATION_PATH) -seq create_$(NAME)_table

migrate-up:
	@migrate -database $(DB_SOURCE) -path $(MIGRATION_PATH) up

migrate-up1:
	@migrate -database $(DB_SOURCE) -path $(MIGRATION_PATH) up 1

migrate-down:
	@migrate -database $(DB_SOURCE) -path $(MIGRATION_PATH) down

migrate-down1:
	@migrate -database $(DB_SOURCE) -path $(MIGRATION_PATH) down 1

migrate-force:
	@migrate -database $(DB_SOURCE) -path $(MIGRATION_PATH) force $(VERSION)

sqlc:
	sqlc generate

test:
	go test -v ./...

server:
	@go run main.go

mock:
	@mockgen -package mockdb -destination db/mock/store.go github.com/anggavb/simplebank/db/sqlc Store
	@echo "Mock generated successfully!"

.PHONY: migrate-create migrate-up migrate-up1 migrate-down migrate-down1 migrate-force sqlc test server mock