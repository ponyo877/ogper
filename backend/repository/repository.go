package repository

import (
	"bytes"
	"context"
	"database/sql"
	_ "embed"
	"strings"
	"text/template"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/ponyo877/ogper/domain"
)

//go:embed ogp.html.tmpl
var ogpHtmlTmpl string

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
	_, err := r.storage.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(r.bucket),
		Key:         aws.String(filename),
		Body:        bytes.NewReader(file),
		ContentType: aws.String(contentType),
	})
	return err
}

func (r *Repository) CreateSite(site *domain.Site) error {
	query := "INSERT INTO sites (hash, title, description, name, site_url, image_url, user_hash) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	if _, err := r.db.Exec(query, site.Hash(), site.Title(), site.Description(), site.Name(), site.SiteURL(), site.ImageURL(), site.UserHash()); err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetSite(hash string) (*domain.Site, error) {
	var title, description, name, siteURL, imageURL, userHash string
	var publishedAt time.Time
	query := "SELECT title, description, name, site_url, image_url, user_hash, published_at FROM sites WHERE hash = $1"
	if err := r.db.QueryRow(query, hash).Scan(&title, &description, &name, &siteURL, &imageURL, &userHash, &publishedAt); err != nil {
		return nil, err
	}
	return domain.NewSite(hash, title, description, name, siteURL, imageURL, userHash, publishedAt), nil
}

func (r *Repository) GetHtml(site *domain.Site) (string, error) {
	tmpl, err := template.New("sample").Parse(ogpHtmlTmpl)
	if err != nil {
		return "", err
	}
	data := map[string]any{
		"title":       site.Title(),
		"description": site.Description(),
		"name":        site.Name(),
		"siteURL":     site.SiteURL(),
		"imageURL":    site.ImageURL(),
	}
	var buf strings.Builder
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (r *Repository) ListSitesByUserID(userHash string) ([]*domain.Site, error) {
	query := "SELECT hash, title, description, name, site_url, image_url, published_at FROM sites WHERE user_hash = $1"
	rows, err := r.db.Query(query, userHash)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sites []*domain.Site
	for rows.Next() {
		var hash, title, description, name, siteURL, imageURL string
		var publishedAt time.Time
		if err := rows.Scan(&hash, &title, &description, &name, &siteURL, &imageURL, &publishedAt); err != nil {
			return nil, err
		}
		site := domain.NewSite(hash, title, description, name, siteURL, imageURL, userHash, publishedAt)
		sites = append(sites, site)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sites, nil
}
