package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ponyo877/repost-ogp-pages/handler"
	mm "github.com/ponyo877/repost-ogp-pages/middleware"
	"github.com/ponyo877/repost-ogp-pages/repository"
	"github.com/ponyo877/repost-ogp-pages/usecase"
)

func main() {
	mux := http.NewServeMux()
	repository := repository.NewRepository()
	usecase := usecase.NewUsecase(repository)
	handler := handler.NewHandler(usecase)
	mux.HandleFunc("POST /ogp", handler.GenerateAltURL)
	log.Printf("running on 8080")
	go http.ListenAndServe(":8080", mm.Logger(mux))

	err := createS3Client()
	if err != nil {
		panic(err)
	}
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/file", postFileHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

func postFileHandler(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	err = uploadFile(file.Filename, src)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "File uploaded")
}

var client *s3.Client

func createS3Client() error {
	hasher := sha256.New()
	hasher.Write([]byte("22df98f6dfd12827a3e0944010bf32d7407f58e3fe2d98f46bbb1eae41f91658"))
	hashedSecretKey := hex.EncodeToString(hasher.Sum(nil))
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("e8dd8f4e091cae4bc66e73b68f76bd47", hashedSecretKey, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		return err
	}

	accountID := os.Getenv("ACCOUNT_ID")

	s3endpoint := fmt.Sprintf("http://%s.r2.cloudflarestorage.com", accountID)

	client = s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(s3endpoint)
	})

	return nil
}

func uploadFile(objectKey string, r io.Reader) error {
	var bucketName string = os.Getenv("AWS_S3_BUCKET_NAME")

	var objectKeyParts []string = strings.Split(objectKey, ".")
	var ext string = "." + objectKeyParts[len(objectKeyParts)-1]
	var contentType string = mime.TypeByExtension(ext)

	_, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(objectKey),
		Body:        r,
		ContentType: aws.String(contentType),
	})

	return err
}
