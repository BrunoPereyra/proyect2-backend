package database

import (
	"backend/config"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GoMongoDB() (*mongo.Database, error) {
	URI := config.URI()
	if URI == "" {
		log.Fatal("MONGODB_URI FATAL")
	}
	clientOptions := options.Client().ApplyURI(URI)

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

func Disconnect(Database *mongo.Database) {
	err := Database.Client().Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
