DB_URL=postgresql://root:123456@localhost:5432/golang_udemy_backend?sslmode=disable

postgres:
	docker run --name golang-udemy-backend --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -d postgres:12-alpine

createdb:
	docker exec -it golang-udemy-backend createdb --username=root --owner=root golang_udemy_backend

dropdb:
	docker exec -it golang-udemy-backend dropdb golang_udemy_backend

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 2

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1


sqlc:
	docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go  github.com/hiamthach/golang-udemy-backend/db/sqlc Store

docker-run:
	docker run --name golang-udemy-backend-docker --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:123456@golang-udemy-backend:5432/golang_udemy_backend?sslmode=disable" golang-udemy-backend:latest

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server migrateup1 migratedown1 test mock docker-run db_docs db_schema proto