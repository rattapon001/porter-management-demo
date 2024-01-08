package job_mongo

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

func LoadMongoDbConfig() *MongoDbConfig {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	return &MongoDbConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
}

func MongoDbInit() *mongo.Client {
	dbConfig := LoadMongoDbConfig()
	uri := "mongodb://" + dbConfig.Username + ":" + dbConfig.Password + "@" + dbConfig.Host + ":" + dbConfig.Port
	println("mongo uri : ", uri)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	return client
}
