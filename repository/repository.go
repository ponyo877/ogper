package repository

import (
	"bytes"
	"context"
	"database/sql"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Repository struct {
	storage *s3.Client
	db      *sql.DB
	bucket  string
}

func NewRepository(storage *s3.Client, db *sql.DB, bucket string) *Repository {
	return &Repository{
		storage,
		db,
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

func (r *Repository) CreateSite(title, description, name, siteURL, imageURL string) error {
	query := "INSERT INTO sites (title, description, name, url, image_url) VALUES (?, ?, ?, ?, ?)"
	if _, err := r.db.Exec(query, title, description, name, siteURL, imageURL); err != nil {
		return err
	}
	return nil
}
