package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "netflix"
const colName = "watchlist"

// Most important

// reference of the mongodb collection
var collection *mongo.Collection

// connect with mongoDB

// it is the initialization method and it runs the first time the application is run and it only runs one time
func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDB
	// context is the go package
	// in mongo.connect function first we need to pass the context type and then the clientOption
	// TODO returns a non-nil, empty Context. Code should use context.TODO when it's unclear which Context to use or it is not yet available (because the surrounding function has not yet been extended to accept a Context parameter).
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")
	collection = client.Database(dbName).Collection(colName)

	// collection instance is ready then
	fmt.Println("Collection instance is ready")
}
