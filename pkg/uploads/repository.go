package uploads

/*
import (
	"context"
	"github.com/rijine/ads-api/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

const (
	COLLECTION = "uploads"
)

type Repository interface {
	Count(countOf string) (int64, error)
}

type repository struct{}

func NewCounterRepository() Repository {
	return &repository{}
}

func (r *repository) UploadImage(countOf string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"countOf": countOf}
	var result Counter
	err := database.Collection(COLLECTION).FindOne(ctx, filter).Decode(&result)

	if err != nil {
		log.Print(err)
		return 0, err
	}

	return result.Count, nil
}
*/
