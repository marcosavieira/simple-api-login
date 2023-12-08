createdb:
	docker exec -it simple-login createdb --username=root --owner=root simple_login

dropdb:
	docker exec -it simple-login dropdb simple_login

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_login?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_login?sslmode=disable" -verbose down

sqlc:
	sqlc generate 	
  

.PHONY: createdb dropdb migrateup migratedown sqlc