include .env

PROJECT ?= digitalbank


DOCKER_COMPOSE_FILE_BUILD=build/docker-compose.yml

DOCKER_COMPOSE_FILE_LOCAL=local/docker-compose.yml

GOLANGCI_LINT_PATH=$$(go env GOPATH)/bin/golangci-lint
GOLANGCI_LINT_VERSION=1.59.0

MIGRATION_FOLDER_PATH=app/gateway/postgres/migrations
GOLANG_MIGRATE_PATH=$$(go env GOPATH)/bin/golang-migrate
GOLANG_MIGRATE_VERSION=4.17.1

start-build:
	docker-compose -f $(DOCKER_COMPOSE_FILE_BUILD) -p $(PROJECT) down -v --remove-orphans
	docker-compose -f $(DOCKER_COMPOSE_FILE_BUILD) -p $(PROJECT) up --build -d --remove-orphans

start-local:
	docker-compose -f $(DOCKER_COMPOSE_FILE_LOCAL) -p $(PROJECT) down --remove-orphans
	docker-compose -f $(DOCKER_COMPOSE_FILE_LOCAL) -p $(PROJECT) up --remove-orphans

migrate-up:
	migrate -path $(MIGRATION_FOLDER_PATH) -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

migrate-down:
	migrate -path $(MIGRATION_FOLDER_PATH) -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down

migrate-new:
	@echo "==> Installing golang-migrate"
	@echo "==> Checking for golang-migrate"
	@if [ -x "$(GOLANG_MIGRATE_PATH)" ]; then \
		echo "Found golang-migrate at $$(which $(GOLANG_MIGRATE_PATH))"; \
		CURRENT_VERSION=$$($(GOLANG_MIGRATE_PATH) --version 2>&1); \
		VERSION_NUM=$$(echo "$$CURRENT_VERSION" | awk '{print $$NF}'); \
		if [ "$$VERSION_NUM" = "$(GOLANG_MIGRATE_VERSION)" ]; then \
			echo "Already installed: $$VERSION_NUM"; \
			exit 0; \
		else \
			echo "Version mismatch, updating..."; \
		fi; \
	else \
		echo "golang-migrate not found, installing..."; \
	fi; \
	if [ "$(OS)" = "linux" ]; then \
		TEMP_DIR=$$(mktemp -d); \
		curl -L https://github.com/golang-migrate/migrate/releases/download/v$(GOLANG_MIGRATE_VERSION)/migrate.linux-amd64.tar.gz | tar xvz -C $$TEMP_DIR; \
		mkdir -p $$(go env GOPATH)/bin; \
		mv $$TEMP_DIR/migrate $(GOLANG_MIGRATE_PATH); \
		chmod +x $(GOLANG_MIGRATE_PATH); \
		rm -rf $$TEMP_DIR; \
	elif [ "$(OS)" = "darwin" ]; then \
		TEMP_DIR=$$(mktemp -d); \
		curl -L https://github.com/golang-migrate/migrate/releases/download/v$(GOLANG_MIGRATE_VERSION)/migrate.darwin-amd64.tar.gz | tar xvz -C $$TEMP_DIR; \
		mkdir -p $$(go env GOPATH)/bin; \
		mv $$TEMP_DIR/migrate $(GOLANG_MIGRATE_PATH); \
		chmod +x $(GOLANG_MIGRATE_PATH); \
		rm -rf $$TEMP_DIR; \
	elif [ "$(OS)" = "windows" ]; then \
		curl -L https://github.com/golang-migrate/migrate/releases/download/v$(GOLANG_MIGRATE_VERSION)/migrate.windows-amd64.zip -o migrate.zip; \
		TEMP_DIR=$$(mktemp -d); \
		unzip migrate.zip -d $$TEMP_DIR; \
		mv $$TEMP_DIR/migrate.exe $$(go env GOPATH)/bin/golang-migrate.exe; \
		rm -rf $$TEMP_DIR migrate.zip; \
	fi
	@echo "==> Creating new migration files for ${name}..."
	$(GOLANG_MIGRATE_PATH) create -ext sql -dir $(MIGRATION_FOLDER_PATH) -seq ${name}

lint:
	@echo "Using golangci-lint path: $(GOLANGCI_LINT_PATH)"
	@echo "Checking golangci-lint version..."
	@$(GOLANGCI_LINT_PATH) --version || echo "golangci-lint not found"
	@echo "==> Installing golangci-lint"
ifeq (,$(shell $(GOLANGCI_LINT_PATH) --version 2>/dev/null | grep $(GOLANGCI_LINT_VERSION)))
	@echo "installing golangci-lint v$(GOLANGCI_LINT_VERSION)"
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v$(GOLANGCI_LINT_VERSION)
else
	@echo "already installed: $(shell eval $(GOLANGCI_LINT_PATH) version)"
endif
	@echo "==> Running golangci-lint"
	@$(GOLANGCI_LINT_PATH) run -c ./.golangci.yml --fix