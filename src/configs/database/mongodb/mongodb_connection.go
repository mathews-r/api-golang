package mongodb

import (
	"context"
	"os"

	"github.com/mathews-r/golang/src/configs/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB_URL  = "DB_URL"
	DB_NAME = "DB_NAME"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongodbURL := os.Getenv(DB_URL)
	dbName := os.Getenv(DB_NAME)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbURL))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("Conexão com o banco de dados realizada com sucesso.")
	return client.Database(dbName), nil
}
