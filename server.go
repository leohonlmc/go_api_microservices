package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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

func respondWithError(c *gin.Context, httpStatusCode int, err error) {
    log.Printf("Error: %v", err)
 
    c.JSON(httpStatusCode, gin.H{"error": err.Error()})
}

func GetAllDocuments(m *MongoDBClient, collectionName string, location, category string) ([]bson.M, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

	collection := m.Client.Database(m.DBName).Collection(collectionName)

	filter := bson.M{}
    if location != "" {
        filter["location"] = location
    }
    if category != "" {
        filter["category"] = category
    }

	cursor, err := collection.Find(ctx, filter)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx) 
	
    var results []bson.M
    if err = cursor.All(ctx, &results); err != nil {
        return nil, err
    }

    return results, nil
}

//get inventories
func getInventoriesHandler(mongoDBClient *MongoDBClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		location := c.Query("location")
		category := c.Query("category")

        documents, err := GetAllDocuments(mongoDBClient, "inventories", location, category)
        if err != nil {
            respondWithError(c, http.StatusInternalServerError, err)
            return
        }
        c.JSON(http.StatusOK, documents)
    }
}

//get news
func getNewsHandler(mongoDBClient *MongoDBClient) gin.HandlerFunc {
	return func(c *gin.Context) {
        documents, err := GetAllDocuments(mongoDBClient, "news", "", "")
        if err != nil {
            respondWithError(c, http.StatusInternalServerError, err)
            return
        }
        c.JSON(http.StatusOK, documents)
    }
}

func helloworld() {
    fmt.Println("Hello World!")
}

func goodbye() {
    fmt.Println("Good Bye!")
}


func main() {
    
    if os.Getenv("ENV") != "production" {
        err := godotenv.Load()
        if err != nil {
            log.Fatal("Error loading .env file")
        }
    }

	dbMongo := os.Getenv("DB_MONGO")
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080" // Default port for local testing
	}

    mongoDBClient, err := NewMongoDBClient(dbMongo, "SFN")
    if err != nil {
        log.Fatal(err)
    }

    cwd, _ := os.Getwd()
	log.Println("Current working directory:", cwd)

    defer mongoDBClient.Close()

	r := gin.Default()

	//get inventories
	r.GET("/inventories", getInventoriesHandler(mongoDBClient))

	//get news
	r.GET("/news", getNewsHandler(mongoDBClient))

	r.Run(":" + port)
}

