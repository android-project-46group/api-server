server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/querier.go github.com/android-project-46group/api-server/db Querier

.PHONY: server mock
