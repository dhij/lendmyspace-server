.PHONY: createdb dropdb migrateup migrateup1 migratedown migratedown1 postgres test server

createdb: 
	docker exec -it lendmyspace_postgres createdb --username=root --owner=root lendmyspace

dropdb: 
	docker exec -it lendmyspace_postgres dropdb lendmyspace

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/lendmyspace?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/lendmyspace?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/lendmyspace?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/lendmyspace?sslmode=disable" -verbose down 1

postgres:
	docker run --name lendmyspace_postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -e POSTGRES_DB=lendmyspace -d postgres:14-alpine

test:
	go test -v -cover ./internal/user/repository ./internal/space/repository

server:
	go run cmd/lendmyspace/main.go