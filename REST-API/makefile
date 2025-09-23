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
	# internal/user/usecases UserService
	mockgen -source=internal/user/usecases/users_usecaseses.go \
	        -destination=$(MOCK_DIR)/user/usecase_mock.go \
	        -package=mocks \
	        $(MODULE)/internal/user/usecases UserService
	# internal/user/repo UserRepository
	mockgen -source=internal/user/repo/users_repo.go \
	        -destination=$(MOCK_DIR)/user/repo_mock.go \
	        -package=mocks \
	        $(MODULE)/internal/user/repo UserRepository
	@echo "✅ Mocks generated in $(MOCK_DIR)"
