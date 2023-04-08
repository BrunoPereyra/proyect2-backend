package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GoMongoDB() *mongo.Database {

	if err := godotenv.Load(); err != nil {
		log.Fatal("ApplyURI db error")
	}
	APPLYURI := os.Getenv("APPLYURI")

	clientOptions := options.Client().ApplyURI(APPLYURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	GoMongoDB := client.Database("goMoongodb")
	if err != nil {
		log.Fatal("error conexion client db")
	}
	return GoMongoDB
}
