package repository

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Repository struct {
	storage *s3.Client
	bucket  string
}

func NewRepository(storage *s3.Client, bucket string) *Repository {
	return &Repository{
		storage,
		bucket,
	}
}

func (r *Repository) GetRedirectPage() error {
	return nil
}

func (r *Repository) PutFile(file []byte, filename, contentType string) error {
	if _, err := r.storage.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(r.bucket),
		Key:         aws.String(filename),
		Body:        bytes.NewReader(file),
		ContentType: aws.String(contentType),
	}); err != nil {
		return err
	}
	return nil
}

func (r *Repository) CreateSite(imageURL, siteURL string) error {
	return nil
}
