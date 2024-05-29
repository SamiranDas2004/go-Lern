package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://samiran4209:Samiran123@cluster0.ungeocm.mongodb.net/"

const dbName = "Netflix"

const colName = "watchList"

// vvi
var collection *mongo.Collection

//connet mongodb

func init() {
	//clint option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongo db
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongoDb connection done")

	collection = client.Database(dbName).Collection(colName)

	//okay

	fmt.Println("collection done")
}
