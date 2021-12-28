.PHONY: createdb dropdb migrateup migratedown

createdb: 
	docker exec -it dplatform_postgres createdb --username=root --owner=root dplatform

dropdb: 
	docker exec -it dplatform_postgres dropdb dplatform

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/dplatform?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/dplatform?sslmode=disable" -verbose down
