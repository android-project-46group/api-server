DB_URL=postgresql://ubuntu:sakamichi@localhost:5432/sakamichi?sslmode=disable

server:
	go run main.go

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path migrations -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path migrations -database "$(DB_URL)" -verbose down 1

populate:
    

mock:
	mockgen -package mockdb -destination db/mock/querier.go github.com/android-project-46group/api-server/db Querier

.PHONY: server postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 mock
