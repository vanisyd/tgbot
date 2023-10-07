package database

import (
	"context"
	"github.com/vanisyd/tgbot/environment"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var DBClient *mongo.Client

type DBCollection interface {
	collectionName() string
}

func Init() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(environment.Env.DBUri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}

	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		log.Fatal(err)
	}

	DBClient = client

	log.Println("[DB] Connected")
}

func getCollection(collection DBCollection) *mongo.Collection {
	return DBClient.Database(environment.Env.DBName).Collection(collection.collectionName())
}
