include .env
export $(shell sed 's/=.*//' .env)

DATABASE_URL ?= $(shell echo $$DATABASE_URL)

migration_up:
	migrate -path internal/db/migrations/ -database "$(DATABASE_URL)" -verbose up

migration_down:
	migrate -path internal/db/migrations/ -database "$(DATABASE_URL)" -verbose down

migration_fix:
	migrate -path internal/db/migrations/ -database "$(DATABASE_URL)" force VERSION

tidy:
	go mod tidy -v
	go fmt ./...

test:
	go test -v ./...

test_with_race:
	go test -v -race ./...