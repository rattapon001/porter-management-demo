package msg_relay

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ChangeCapture(client *mongo.Client, dataBase string, collectionName string, ctx context.Context) *mongo.ChangeStream {
	collection := client.Database(dataBase).Collection(collectionName)
	pipeline := []bson.M{
		{"$match": bson.M{"operationType": bson.M{"$in": []string{"insert", "update"}}}},
	}
	changeStream, err := collection.Watch(ctx, pipeline, options.ChangeStream())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Change Data Capture started")
	return changeStream
}
