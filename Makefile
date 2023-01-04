postgres:
	docker run --name postgres-db -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secretpw -d postgres

createdb:
	docker exec -it postgres-db createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgres-db dropdb simplebank

migrateup:
	migrate -path db/migrations -database "postgresql://simplebank_backend:456fgrrg@localhost:5432/simple_bank?sslmode=disable"  -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://simplebank_backend:456fgrrg@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown