package database

import (
	"backend/config"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Pool *mongo.Client
}

func NewMongoDB(poolSize uint64) (*MongoDB, error) {
	// se crea la configuración del cliente con la URI y el tamaño de la pool
	URI := config.URI()
	if URI == "" {
		log.Fatal("MONGODB_URI FATAL")
	}
	clientOptions := options.Client().ApplyURI(URI).SetMaxPoolSize(poolSize)

	// se crea la pool de conexiones
	pool, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	// se establece la conexión a la base de datos
	if err = pool.Connect(context.Background()); err != nil {
		return nil, err
	}

	// se verifica que la conexión a la base de datos sea exitosa.
	if err = pool.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	return &MongoDB{Pool: pool}, nil
}
