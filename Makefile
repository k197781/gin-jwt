.PHONY: test
test:
	go run db/migrate/migrate.go
	go test ./... -v