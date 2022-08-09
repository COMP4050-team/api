.PHONY: generate test coverage

generate:
	go generate ./...

test:
	go test -v -coverprofile=cover.out ./...

coverage:
	go tool cover -html=cover.out

run:
	go run ./...
