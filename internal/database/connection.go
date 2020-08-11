package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	timeout             = 2
	ErrDatabaseConn     = "Database connection failed"
	errNilClientMsg     = "Client must not be nil"
	errEmptyDbNameMsg   = "Database name must not be empty"
	errEmptyCollNameMsg = "Collection name must not be empty"
)

type mongodb struct {
	Client   *mongo.Client
	Database *mongo.Database
}

var Mongo mongodb

func getContext() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	return ctx
}

func Connect() error {
	dbOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(getContext(), dbOptions)
	if err != nil {
		return err
	}

	Mongo = mongodb{
		Client:   client,
		Database: client.Database("ads"),
	}

	return nil
}
