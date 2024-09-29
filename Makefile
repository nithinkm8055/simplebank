.PHONY: postgres createdb migrateup migratedown sqlc test

postgres:
	docker run -d -p 5432:5432 --name postgres16 -e POSTGRES_PASSWORD=secret  -e POSTGRES_USER=root postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb -U root --owner=root simple_bank

dropdb:
	docker exec -it postgres16 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" --verbose down

sqlc:
	sqlc generate

test:
	go test -cover -v ./...
