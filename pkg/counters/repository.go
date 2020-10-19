package counters

import (
	"context"
	"time"

	"github.com/rijine/ads-api/internal/database"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	COLLECTION = "counters"
)

type Repository interface {
	GetAndUpdate(countOf string) (int64, error)
}

type repository struct{}

func NewCounterRepository() Repository {
	return &repository{}
}

func (r *repository) GetAndUpdate(countOf string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"countOf": countOf}
	var result Counter
	err := database.Collection(COLLECTION).FindOne(ctx, filter).Decode(&result)
	if err.Error() == "mongo: no documents in result" {

		insrt := Counter{
			CountOf: "user",
			Count:   1,
		}

		_, err = database.Collection(COLLECTION).InsertOne(ctx, insrt)

	}
	if err != nil {

		return 0, err
	}
	update := bson.D{{"$set", bson.M{"count": result.Count + 1}}}

	_, err = database.Collection(COLLECTION).UpdateOne(ctx, filter, update)

	if err != nil {
		return 0, err
	}

	return result.Count, nil
}
