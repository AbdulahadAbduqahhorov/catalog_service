run: 
	go run cmd/main.go
migrateup:
	migrate -path ./migrations/postgres -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/catalog_service_db?sslmode=disable' up
migratedown:
	migrate -path ./migrations/postgres -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/catalog_service_db?sslmode=disable' down
pull-proto-module:
	git submodule update --init --recursive

update-proto-module:
	git submodule update --remote --merge