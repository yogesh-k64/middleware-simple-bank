test: 
	echo hello

postgres:
	docker run --name postgres17 --network bank-network -e POSTGRES_USER=root -p 5432:5432 -e POSTGRES_PASSWORD=1234 -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --user=root --owner=root simple_bank

dropdb:
	docker exec -it postgres17 dropdb simple_bank

migrateup:
	migrate -path db/migration/ -database "postgresql://root:1234@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration/ -database "postgresql://root:1234@localhost:5432/simple_bank?sslmode=disable" -verbose down

migrateup1:
	migrate -path db/migration/ -database "postgresql://root:1234@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path db/migration/ -database "postgresql://root:1234@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	 mockgen -package mockdb -destination db/mock/store.go github.com/yogesh-k64/middleware-simple-bank/db/sqlc Store

.PHONY: test postgres createdb dropdb migrateup migratedown sqlc mock migratedown1 migrateup1 server