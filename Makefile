up:
	@docker-compose up -d

down:
	@docker-compose down

migrate: migrateBuild
	@./bin/migrate $(filter-out $@,$(MAKECMDGOALS))

migrateBuild:
	@go build -o bin/migrate cmd/migrate/main.go

%:
	@: