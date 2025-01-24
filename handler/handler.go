package handler

import (
	"io"
)

type Handler struct {
	usecase Usecase
}

func NewHandler(usecase Usecase) *Handler {
	return &Handler{usecase: usecase}
}

type Usecase interface {
	GenerateOGPPage(title, description, name, siteURL string, file io.Reader, size int64) (string, error)
	GetOGPPage(hash string) (string, error)
}
