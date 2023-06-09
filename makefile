.PHONY: injectapp
injectapp:
	cd ./internal/app && wire

.PHONY: runapp
runapp:
	go run ./cmd/app/main.go

.PHONY: documentation
docs:
	swag init -g ./cmd/app/main.go -o ./docs

.PHONY: migrateup
migrateup:
	go run db/migration/main.go -migration=up

.PHONY: migratedown
migratedown:
	go run db/migration/main.go -migration=down