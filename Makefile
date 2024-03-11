postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres16 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

makeFileDir := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
sqlc1:
	docker run --rm -v $(makeFileDir):/src -w /src kjconroy/sqlc generate

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

start:
	docker start postgres16

stop:
	docker stop postgres16

list:
	docker ps

check:
	docker images 
	
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test start stop list check