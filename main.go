package main

import (
	"log"
	"net/http"

	"github.com/ponyo877/repost-ogp-pages/config"
	"github.com/ponyo877/repost-ogp-pages/handler"
	"github.com/ponyo877/repost-ogp-pages/middleware"
	"github.com/ponyo877/repost-ogp-pages/repository"
	"github.com/ponyo877/repost-ogp-pages/usecase"
)

func main() {
	mux := http.NewServeMux()
	storage, bucket, err := config.NewCloudflareR2Config()
	if err != nil {
		log.Fatal(err)
	}
	repository := repository.NewRepository(storage, bucket)
	usecase := usecase.NewUsecase(repository)
	handler := handler.NewHandler(usecase)
	mux.HandleFunc("POST /upload", handler.GenerateAltURL)
	log.Printf("running on 8080")
	http.ListenAndServe(":8080", middleware.Logger(mux))
}
