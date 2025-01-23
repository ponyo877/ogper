package handler

import "io"

type Handler struct {
	usecase Usecase
}

func NewHandler(usecase Usecase) *Handler {
	return &Handler{usecase: usecase}
}

type Usecase interface {
	GenerateAltURL(title, description, name, siteURL string, file io.Reader, size int64) error
}
