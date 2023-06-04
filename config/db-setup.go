package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

func initDb() *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://g6-1c-2023:7qrq4AI2GGXBqvVK@recommendationdb.jop2vfk.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	return client
}

func setRules(collection *mongo.Collection) error {
	count, err := collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		fmt.Print("Error 2\n")
		return err
	}
	if count == 0 {
		proximityDoc := map[string]interface{}{
			"rule":    "proximity",
			"radius":  20,
			"enabled": true,
		}
		interestsDoc := map[string]interface{}{
			"rule":    "interests",
			"enabled": true,
		}
		documents := []interface{}{
			proximityDoc,
			interestsDoc,
		}
		_, err := collection.InsertMany(ctx, documents)
		if err != nil {
			fmt.Print("Error 3\n")
			return err
		}
	}
	return nil
}

func SetUpDB() (*mongo.Client, error) {
	client := initDb()
	err := client.Ping(ctx, nil)
	if err != nil {
		fmt.Print("Error 1\n")
		return nil, err
	}
	collection := client.Database("recDB").Collection("rules")
	err = setRules(collection)
	if err != nil {
		fmt.Print("Error 4\n")
		return nil, err
	}
	return client, nil
}

func TearDownDB(client *mongo.Client) {
	err := client.Disconnect(ctx)
	if err != nil {
		fmt.Print("Error 5\n")
		panic(err)
	}
}
