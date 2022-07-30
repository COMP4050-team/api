package main

//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -source=internal/pkg/db/db.go -destination=fixtures/mocks/mocks.go -package=mocks
//go:generate go run github.com/99designs/gqlgen generate
