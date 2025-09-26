migrate-up:
	@echo "Checking if migrations need to run..."
	@if docker exec -i $(DB_CONTAINER) psql -U $$POSTGRES_USER -d $$POSTGRES_DB \
	    -c "SELECT 1 FROM schema_version LIMIT 1;" >/dev/null 2>&1; then \
	    echo "Database already migrated. Skipping."; \
	else \
	    echo "Running migrations..."; \
	    cd db/migrations;tern migrate; \
	fi


generate-query:
	cd db; sqlc generate;

MODULE := rest-api
# directories for output
MOCK_DIR := internal/mocks


mocks:
	@echo "Generating GoMock mocks..."
	# db/sqlc Querier
	mockgen -source=db/sqlc/querier.go \
	        -destination=$(MOCK_DIR)/sqlc/querier_mock.go \
	        -package=mocks \
	        $(MODULE)/db/sqlc querier.go
	# internal/user/service UserService
	mockgen -source=internal/user/service/users_service.go \
	        -destination=$(MOCK_DIR)/user/user_service_mock.go \
	        -package=mocks \
	        $(MODULE)/internal/user/service UserService
	# internal/user/repo UserRepository
	mockgen -source=internal/user/repo/users_repo.go \
	        -destination=$(MOCK_DIR)/user/repo_mock.go \
	        -package=mocks \
	        $(MODULE)/internal/user/repo UserRepository
	@echo "✅ Mocks generated in $(MOCK_DIR)"


APP_NAME := rest-api
BUILD_DIR := build
MAIN_PKG := ./cmd/server   # adjust if needed

.PHONY: build clean

build:
	GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PKG)

server: build
	@env $$(grep -v '^#' .env | xargs) ./$(BUILD_DIR)/$(APP_NAME)

clean-build:
	rm -rf $(BUILD_DIR)

image:
	docker build -t $(APP_NAME):1.0.0 .

image-run: image
	docker run --name rest-app -p 8080:8081\
	  --env-file .env \
	 $(APP_NAME):1.0.0
.PHONY: db
DB_CONTAINER:= rest-db
db:
	@if [ "$$(docker ps -q -f name=$(DB_CONTAINER))" ]; then \
		echo "Database container '$(DB_CONTAINER)' is already running."; \
	else \
		echo "Starting database container..."; \
		docker compose up -d db; \
		echo "Database container started."; \
	fi

db-migrate: db
	@echo "Waiting 2 seconds for database to initialize..."
	sleep 2
	$(MAKE) migrate-up

build-and-run-app: db-migrate

	@echo "starting the rest-api application"
	docker compose up -d app
	@echo "application is started"
	@echo "Checking application health..."
	curl --fail http://localhost:8080/healthcheck || \
		(echo "Healthcheck failed!" && exit 1)
