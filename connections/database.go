package connections

import (
	"context"
	"example/go-crud/config"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(config.EnvMongoURI()).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// Client instance
var db = ConnectDB()

// getting database collections
func GetCollection(collectionName string) *mongo.Collection {
	collection := db.Database(os.Getenv("MONGO_DB")).Collection(collectionName)
	return collection
}
