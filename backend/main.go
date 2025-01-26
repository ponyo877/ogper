package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ponyo877/ogper/config"
	"github.com/ponyo877/ogper/handler"
	"github.com/ponyo877/ogper/middleware"
	"github.com/ponyo877/ogper/repository"
	"github.com/ponyo877/ogper/usecase"
)

var healthCheckUrl = os.Getenv("HEALTH_CHECK_URL")

func main() {
	mux := http.NewServeMux()
	storage, bucket, err := config.NewCloudflareR2Config()
	if err != nil {
		log.Fatal(err)
	}
	db, err := config.NewPostgreSQLConfig()
	if err != nil {
		log.Fatal(err)
	}
	repository := repository.NewRepository(storage, db, bucket)
	usecase := usecase.NewUsecase(repository)
	handler := handler.NewHandler(usecase)
	mux.HandleFunc("POST /upload", handler.GenerateOGPPage)
	mux.HandleFunc("GET /{hash}", handler.GetOGPPage)
	mux.HandleFunc("GET /links", handler.ListSitesByUserID)
	log.Printf("running on 8080")
	// for render sleep
	go func() {
		for {
			http.Get(healthCheckUrl)
			time.Sleep(10 * time.Minute)
		}
	}()
	http.ListenAndServe(":8080", middleware.CORS(middleware.Logger(mux)))
}
