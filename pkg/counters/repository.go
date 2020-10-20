package counters

import (
	"context"
	"fmt"
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

	filter := bson.M{"countof": countOf}
	var result Counter
	err := database.Collection(COLLECTION).FindOne(ctx, filter).Decode(&result)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {

			insrt := Counter{
				CountOf: countOf,
				Count:   1,
			}

			_, err = database.Collection(COLLECTION).InsertOne(ctx, insrt)
			if err != nil {
				fmt.Println("Counter error:", err)
				return 0, err
			}
			return 1, nil
		}
		return 0, err
	}

	ct1 := result.Count + 1

	update := bson.M{"$set": bson.M{"count": ct1}}

	_, err = database.Collection(COLLECTION).UpdateOne(ctx, filter, update)

	if err != nil {
		fmt.Println("update", err)
		return 0, err
	}

	return result.Count, nil
}
