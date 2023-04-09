package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GoMongoDB() (*mongo.Database, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("ApplyURI db error")
	}
	APPLYURI := os.Getenv("APPLYURI")

	clientOptions := options.Client().ApplyURI(APPLYURI)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	// Se establece una conexión a la base de datos.
	if err = client.Connect(context.Background()); err != nil {
		return nil, err
	}

	// Se verifica que la conexión a la base de datos sea exitosa.
	if err = client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	return client.Database("goMoongodb"), nil
}
