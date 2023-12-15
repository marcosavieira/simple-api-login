postgres:
	docker run --name simple-login -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret d postgres:16-alpine

createdb:
	docker exec -it simple-login createdb --username=root --owner=root simple_login

dropdb:
	docker exec -it simple-login dropdb simple_login

migrateup:
	migrate -path db/migration -database "postgresql://postgres:09021328@db-simple-login.co2yqtstfx61.sa-east-1.rds.amazonaws.com?sslmode=require" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_login?sslmode=require" -verbose down

sqlc:
	sqlc generate 	
  

.PHONY: postgres createdb dropdb migrateup migratedown sqlc