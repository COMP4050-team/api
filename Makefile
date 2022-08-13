.PHONY: generate test coverage

generate:
	go generate ./...

test:
	go test -v -coverprofile=cover.out ./...

coverage:
	go tool cover -html=cover.out

run:
	JWT_SECRET=catjam DB_FILE_PATH=test.db TEST_EXECUTOR_ENDPOINT=http://127.0.0.1:8080/ go run ./...
