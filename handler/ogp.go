package handler

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (h *Handler) GenerateAltURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		log.Printf("Error : %v\n", err)
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	// Retrieve the file from form data
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error : %v\n", err)
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Retrieve the url from form data
	url := r.FormValue("url")
	if url == "" {
		log.Printf("Error : %v\n", err)
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	// ファイルサイズを取得
	fileSize := handler.Size

	// 署名付きペイロードでアップロード
	if err := UploadFileWithSignedPayload(handler.Filename, file, fileSize); err != nil {
		log.Printf("Error : %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s\n", handler.Filename)
	fmt.Fprintf(w, "URL: %s\n", url)
}

func (h *Handler) GetRedirectPage(w http.ResponseWriter, r *http.Request) {
	if err := h.usecase.GetRedirectPage(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Hello, World!\n"))
}

func UploadFileWithSignedPayload(filename string, f io.Reader, size int64) error {
	accountID := os.Getenv("ACCOUNT_ID")
	accessKeyID := os.Getenv("ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("ACCESS_KEY_SECRET")

	endpoint := fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountID)
	bucket := "ogp"

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, accessKeySecret, "")),
		config.WithRegion("auto"),
		// config.WithRequestChecksumCalculation(0),
		// config.WithResponseChecksumValidation(0),
	)
	if err != nil {
		log.Printf("failed to LoadDefaultConfig: %v", err)
		return fmt.Errorf("failed to load config: %w", err)
	}

	client := s3.NewFromConfig(
		cfg,
		func(o *s3.Options) {
			o.BaseEndpoint = aws.String(endpoint)
			// o.UseAccelerate = false
			// o.UsePathStyle = true
		},
	)

	_, err = client.PutObject(
		context.TODO(),
		&s3.PutObjectInput{
			Bucket:        aws.String(bucket),
			Key:           aws.String(filename),
			Body:          f,
			ContentLength: aws.Int64(size),
			ContentType:   aws.String("image/png"),
		},
	)
	if err != nil {
		log.Printf("failed to PutObject: %s", err.Error())
		return fmt.Errorf("failed to upload file: %w", err)
	}
	// presignClient := s3.NewPresignClient(client)

	// presignResult, err := presignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
	// 	Bucket: aws.String(bucket),
	// 	Key:    aws.String("images.jpeg"),
	// })

	// if err != nil {
	// 	panic("Couldn't get presigned URL for PutObject")
	// }

	// fmt.Printf("Presigned URL For object: %s\n", presignResult.URL)
	return nil
}
