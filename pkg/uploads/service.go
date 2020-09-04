package uploads

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/dgryski/trifles/uuid"
	"google.golang.org/api/option"
	"io"
)

type Service interface {
	//UploadImage(ctx context.Context, r io.Reader, kind *string) (string, error)
	UploadImage(ctx context.Context, picture graphql.Upload, kind *string) (string, error)
}

type service struct{}

const publicURL = "https://storage.googleapis.com/%s/%s"

var (
	storageClient *storage.Client
	options       = option.WithCredentialsFile("configs/gcloud.keys.json")
)

func NewCounterService() Service {
	return &service{}
}

func (s *service) UploadImage(ctx context.Context, picture graphql.Upload, kind *string) (string, error) {
	bucket := "imerch-images"
	//bucket := "posts"
	var err error
	storageClient, err = storage.NewClient(ctx, options)
	if err != nil {
		return "", err
	}

	filename := *kind + "/" + uuid.UUIDv4() + "-" + picture.Filename
	sw := storageClient.Bucket(bucket).Object(filename).NewWriter(ctx)

	if _, err := io.Copy(sw, picture.File); err != nil {
		return "", err
	}

	if err := sw.Close(); err != nil {
		return "", err
	}

	return fmt.Sprintf(publicURL, bucket, filename), nil
}
