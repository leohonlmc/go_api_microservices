package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joho/godotenv"
)

type MongoDBClient struct {
    Client *mongo.Client
    DBName string
}

func NewMongoDBClient(uri, dbName string) (*MongoDBClient, error) {
    clientOptions := options.Client().ApplyURI(uri)
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        return nil, err
    }

    return &MongoDBClient{
        Client: client,
        DBName: dbName,
    }, nil
}

func (m *MongoDBClient) Close() error {
    return m.Client.Disconnect(context.Background())
}

func GetAllDocuments(m *MongoDBClient, collectionName string) ([]bson.M, error) {
	collection := m.Client.Database(m.DBName).Collection(collectionName)

    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    var results []bson.M
    if err = cursor.All(context.Background(), &results); err != nil {
        return nil, err
    }

    return results, nil
}

//get inventories
func getInventoriesHandler(mongoDBClient *MongoDBClient) gin.HandlerFunc {
	return func(c *gin.Context) {
        documents, err := GetAllDocuments(mongoDBClient, "inventories")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, documents)
    }
}

func main() {
	err := godotenv.Load()

	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	dbMongo := os.Getenv("DB_MONGO")

    mongoDBClient, err := NewMongoDBClient(dbMongo, "SFN")
    if err != nil {
        log.Fatal(err)
    }
    defer mongoDBClient.Close()

	r := gin.Default()

	r.GET("/inventories", getInventoriesHandler(mongoDBClient))

	r.Run(":8080")
}

