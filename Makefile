DB_URL=postgresql://ubuntu:sakamichi@localhost:5432/sakamichi?sslmode=disable

server:
	go run main.go

migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down

mock:
	mockgen -package mockdb -destination db/mock/querier.go github.com/android-project-46group/api-server/db Querier

psql:
	psql -h localhost -p 5432 -U ubuntu -d sakamichi

populate:
	go run db/data/populate/main.go

test:
	go test -cover -shuffle=on ./api

.PHONY: server migrateup migratedown mock psql populate test
