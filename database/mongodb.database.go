package database

import (
	"context"
	"crypto/tls"
	"log"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/heyubani/go-template/config"
)

var mongoClient *mongo.Client

const DB_CONTEXT_TIMEOUT time.Duration = 90 * time.Second

func GetMongoDbConnection() *mongo.Client {
	return mongoClient
}

// For unit test cases, we might need a mock connection to the database, this method does set the mock connection
// on the file
func SetDbConnection(db *mongo.Client) {
	mongoClient = db
}

func ConnectMongo() func() {
	if mongoClient == nil {
		tlsConfig := &tls.Config{}
		tlsConfig.InsecureSkipVerify = true
		uri := strings.TrimSpace(config.AppConfig.DbUrl)
		env := strings.TrimSpace(config.AppConfig.Env)
		log.Println("Mongo URI: ", env)

		var connectOptions *options.ClientOptions
		if strings.ToLower(env) == "development" {
			connectOptions = options.Client().ApplyURI(uri)
		} else {
			connectOptions = options.Client().ApplyURI(uri).SetTLSConfig(tlsConfig)
		}

		ctx, cancel := context.WithTimeout(context.Background(), DB_CONTEXT_TIMEOUT)
		defer cancel()
		client, err := mongo.Connect(ctx, connectOptions)

		if err != nil {
			panic(err)
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			panic(err)
		}

		log.Print("Mongo Db connected  ✓")
		mongoClient = client
	}

	return func(client *mongo.Client) func() {
		return func() {
			if err := client.Disconnect(context.Background()); err != nil {
				panic(err)
			}

			mongoClient = nil
		}
	}(mongoClient)
}

func GetMongoCollection(dBName string, dbCollection string) *mongo.Collection {
	collection := mongoClient.Database(dBName).Collection(dbCollection)
	return collection
}

func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
}
