.PHONY: help
help:
	go run cmd/*.go

.PHONY: server
server:
	go run cmd/*.go server

.PHONY: migrate
migrate:
	BUNDEBUG=2 go run cmd/*.go db migrate

.PHONY: create_migration
create_migration:
	go run cmd/*.go db create ${name}

.PHONY: lint
lint:
	golangci-lint run -c ./golangci.yaml -v

.PHONY: fix_lint
fix_lint:
	golangci-lint run -c ./golangci.yaml -v --fix

.PHONY: gen_docs
gen_docs:
	swag init -g ./cmd/main.go -o ./docs --parseDependency --parseInternal --quiet
