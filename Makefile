.PHONY: createdb dropdb migrateup migratedown postgres test server

createdb: 
	docker exec -it dplatform_postgres createdb --username=root --owner=root lendmyspace

dropdb: 
	docker exec -it dplatform_postgres dropdb lendmyspace

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/lendmyspace?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/lendmyspace?sslmode=disable" -verbose down

postgres:
	docker run --name dplatform_postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -e POSTGRES_DB=dplatform -d postgres:14-alpine

test:
	go test -v -cover ./internal/user/repository ./internal/space/repository

server:
	go run cmd/lendmyspace/main.go