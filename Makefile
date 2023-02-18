DB_URL=postgresql://ubuntu:sakamichi@localhost:5432/sakamichi?sslmode=disable

.DEFAULT_GOAL := help

server:
	go run main.go

migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down

migrateclear:	## migrate に失敗したものを更新する。
	psql -h localhost -p 5432 -U ubuntu -d sakamichi -c "UPDATE schema_migrations SET dirty='f';"

mock:	## mock を生成する
	mockgen -package mockdb -destination db/mock/querier.go github.com/android-project-46group/api-server/db Querier

psql:
	psql -h localhost -p 5432 -U ubuntu -d sakamichi

populate:	## データを投入する
	go run db/data/populate/main.go

sqlboiler:	## DB 情報を元に SQL Boiler の構造体を再生成する
	cd ./db && sqlboiler psql

help:	## https://postd.cc/auto-documented-makefile/
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

godoc:	## godoc をローカルで表示する。http://localhost:8080/{module_name}
	pkgsite

test:
	go test -cover -shuffle=on ./api

.PHONY: server migrateup migratedown mock psql populate test
.PHONY: sqlboiler help godoc
