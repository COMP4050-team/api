.PHONY: generate test coverage

generate:
	go generate ./...

test:
	go test -coverprofile=cover.out ./...

coverage:
	go tool cover -html=cover.out
