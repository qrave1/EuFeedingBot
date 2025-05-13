migrate:
	@goose sqlite3 application.db -dir ./cmd/migrations up
