package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//GetMongoDbConnection get connection of mongodb
func GetMongoDbConnection() (*mongo.Client, error) {
	client, err := mongo.Connect(
		context.Background(), options.Client().ApplyURI(
			CONNECTION_URI+CONNECTION_PORT))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	return client, nil
}

func GetMongoDbCollection(DbName string, CollectionName string) (*mongo.Collection, *mongo.Client, error) {
	client, err := GetMongoDbConnection()

	if err != nil {
		return nil, nil, err
	}

	collection := client.Database(DbName).Collection(CollectionName)

	return collection, client, nil
}
