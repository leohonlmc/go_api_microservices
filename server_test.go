package main

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

//unit test for NewMongoDBClient
func TestNewMongoDBClient(t *testing.T) {
    err := godotenv.Load()
    if err != nil {
        t.Fatalf("Error loading .env file: %v", err)
    }

    dbMongo := os.Getenv("DB_MONGO")
    client, err := NewMongoDBClient(dbMongo, "SFN")
    if err != nil {
        t.Errorf("NewMongoDBClient returned an error: %v", err)
    }
    if client == nil {
        t.Errorf("NewMongoDBClient returned a nil client")
    }

    if client != nil {
        if err := client.Close(); err != nil {
            t.Errorf("Failed to close MongoDB client: %v", err)
        }
    }
}



