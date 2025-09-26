db-migrate:
	cd db/migrations;tern migrate;

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
