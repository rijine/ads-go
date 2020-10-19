package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	timeout             = 5
	ErrDatabaseConn     = "Database connection failed"
	errNilClientMsg     = "Client must not be nil"
	errEmptyDbNameMsg   = "Database name must not be empty"
	errEmptyCollNameMsg = "Collection name must not be empty"
)

type mongodb struct {
	Client   *mongo.Client
	Database *mongo.Database
}

var db mongodb

// TODO: No effect
func NewContext() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	return ctx
}

func Collection(name string) *mongo.Collection {
	return db.Database.Collection(name)
}

func Connect() error {
	dbOptions := options.Client().ApplyURI("mongodb+srv://dbadmin:admin123@mymongocluster.1p40q.mongodb.net/v-ads?retryWrites=true&w=majority")
	client, err := mongo.Connect(NewContext(), dbOptions)
	if err != nil {
		return err
	}

	db = mongodb{
		Client:   client,
		Database: client.Database("v-ads"),
	}

	return nil
}
