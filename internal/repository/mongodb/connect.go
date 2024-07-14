package mongo

import (
	"log"
	"os"

	"github.com/ThembinkosiThemba/go-project-starter/pkg/http"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(dbName string) (*mongo.Database, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	// Remember to create an env file and put your variables
	mongoUri := os.Getenv("MONGO_URL")

	var ctx, cancel = http.Context()
	defer cancel()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoUri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("Connected to Mongo DB")

	db := client.Database(dbName)
	return db, nil
}
