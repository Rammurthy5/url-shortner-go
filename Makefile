include .env
export $(shell sed 's/=.*//' .env)

DATABASE_URL ?= $(shell echo $$DATABASE_URL)

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: run quality control checks
.PHONY: audit
#audit: test
audit:
	go mod tidy -diff
	go mod verify
	test -z "$(shell gofmt -l .)"
	go vet ./...
	golangci-lint run ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## test: run all tests
.PHONY: test
test:
	#go test -v -race -buildvcs ./...
	go test -v ./...

## pre-commit: run pre-commit checks
.PHONY: pre-commit
pre-commit: audit

## build: Build the binary
.PHONY: build
build:
	go build -o bin/url-shortener cmd/web-app/main.go

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## format all
.PHONY: format
fmt:
	gofmt -w -s -d .

# ==================================================================================== #
# DATABASE MIGRATIONS
# ==================================================================================== #

migration_up:
	migrate -path internal/db/migrations/ -database "$(DATABASE_URL)" -verbose up

migration_down:
	migrate -path internal/db/migrations/ -database "$(DATABASE_URL)" -verbose down

migration_fix:
	migrate -path internal/db/migrations/ -database "$(DATABASE_URL)" force VERSION
