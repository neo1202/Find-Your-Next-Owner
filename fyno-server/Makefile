ifneq ("$(wildcard .env)", "")
    include .env
endif

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm 
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

## run: run the app
.PHONY: run
run:
	go run ./cmd -db-dsn=${DB_DSN}

# test: run all tests
.PHONY: test
test:
	go test -v ./...

## db/psql: connect to the database using psql
.PHONY: db/psql 
db/psql:
	psql ${DB_DSN}
	
## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new 
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./db/migrations ${name}
	
## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up 
db/migrations/up: confirm
	@echo 'Running up migrations...'
	migrate -path ./db/migrations -database ${DB_DSN} up

## db/migrations/down: apply all down database migrations
.PHONY: db/migrations/down
db/migrations/down: confirm
	@echo 'Running down migrations...'
	migrate -path ./db/migrations -database ${DB_DSN} down