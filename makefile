postgres:
	docker run --name postgres12 -p 5433:5432 -e POSTGRES_USER="root" -e POSTGRES_PASSWORD=07May2002@ -e POSTGRES_DB="root" -d postgres:16.14-alpine3.23
 
createdb: 
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:07May2002@@localhost:5433/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:07May2002@@localhost:5433/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

mock:
	mockgen -destination db/mock/store.go -package mockdb github.com/quavo19/bankapi/db/sqlc Store

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc mock test server