postgres:
	docker run --name golang-udemy-backend --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -d postgres:12-alpine

createdb:
	docker exec -it golang-udemy-backend createdb --username=root --owner=root golang_udemy_backend

dropdb:
	docker exec -it golang-udemy-backend dropdb golang_udemy_backend

migrateup:
	migrate -path db/migration -database "postgresql://root:123456@localhost:5432/golang_udemy_backend?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:123456@localhost:5432/golang_udemy_backend?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:123456@localhost:5432/golang_udemy_backend?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:123456@localhost:5432/golang_udemy_backend?sslmode=disable" -verbose down 1


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

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server migrateup1 migratedown1 test mock