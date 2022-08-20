package config

import (
	"flag"
	"log"
	"os"
)

type Config struct {
	Port                 int
	JWTSecret            string
	DBFilePath           string
	TestExecutorEndpoint string
}

func NewConfig() Config {
	c := Config{}

	flag.StringVar(&c.JWTSecret, "jwt-secret", os.Getenv("JWT_SECRET"), "The JWT secret to use. Required")
	flag.IntVar(&c.Port, "port", 8080, "The port to listen on. Default is 8080")
	flag.StringVar(&c.DBFilePath, "db-path", "test.db", "The path to the sqlite3 database. Default is test.db")
	flag.StringVar(&c.TestExecutorEndpoint, "test-executor-endpoint", "http://localhost:8080/", "The endpoint to the test executor. Default is http://localhost:8080/")

	flag.Parse()

	if c.JWTSecret == "" {
		log.Fatal("The JWT secret is required")
	}

	return c
}
