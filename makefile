up:
	go run main.go

mock:
	mockgen -source=database/db_interface.go -destination=database/mocks/db_mock.go DB -package=mocks

test: |
	make mock
	go test ./...