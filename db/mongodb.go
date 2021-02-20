package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type _MongoDBClient struct {
	client  *mongo.Client
	context context.Context
	dbname  string
}

var _mongo *_MongoDBClient

// Connect to a MongoDB Instance
func Connect(uri string, dbname string) {
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	ctxPing, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(ctxPing, readpref.Primary())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	log.Print("Connected to MongoDB Server : ", uri)
	_mongo = &_MongoDBClient{
		client,
		ctx,
		dbname,
	}
}

// GetCollection returns a MongoDB collection
func GetCollection(collection string) *mongo.Collection {
	return _mongo.client.Database(_mongo.dbname).Collection(collection)
}

// Close the connexion to the DB
func Close() {
	err := _mongo.client.Disconnect(_mongo.context)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// log.Print(fmt.Sprintf("Connection to MongoDB Server [%s/%s] closed.", _mongo.client. ClientOptions().uri, _mongo.dbname))
	log.Print(fmt.Sprintf("Connection to MongoDB Server [%s] closed.", _mongo.dbname))
}
