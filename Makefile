up:
	@docker-compose up -d

down:
	@docker-compose down

build:
	@docker-compose build

migrate: migrateBuild
	@./bin/migrate $(filter-out $@,$(MAKECMDGOALS))

migrateBuild:
	@go build -o bin/migrate cmd/migrate/main.go

test:
	@go test -v ./...

%:
	@: