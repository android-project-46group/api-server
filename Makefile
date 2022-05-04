DB_URL=postgresql://ubuntu:sakamichi@localhost:5432/sakamichi?sslmode=disable

server:
	go run main.go

migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down

mock:
	mockgen -package mockdb -destination db/mock/querier.go github.com/android-project-46group/api-server/db Querier

populate:
	go run db/data/populate/main.go

.PHONY: server migrateup migratedown mock populate
