test: 
	echo hello

postgres:
	docker run --name postgres17 -e POSTGRES_USER=root -p 5432:5432 -e POSTGRES_PASSWORD=1234 -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --user=root --owner=root simple_bank

dropdb:
	docker exec -it postgres17 dropdb simple_bank

migrateup:
	migrate -path db/migration/ -database "postgresql://root:1234@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration/ -database "postgresql://root:1234@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: test postgres createdb dropdb migrateup migratedown sqlc server