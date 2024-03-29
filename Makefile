postgres:
	docker run --name postgres-db -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=456fGrrg -d postgres

createdb:
	docker exec -it postgres-db createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgres-db dropdb simple_bank

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:456fGrrg@localhost:5432/simple_bank?sslmode=disable"  -verbose up

migrateup1:
	migrate -path db/migrations -database "postgresql://postgres:456fGrrg@localhost:5432/simple_bank?sslmode=disable"  -verbose up 1

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:456fGrrg@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migrations -database "postgresql://postgres:456fGrrg@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/judgegodwins/simplebank/db/sqlc Store

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 test server mock