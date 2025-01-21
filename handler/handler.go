package handler

import "io"

type Handler struct {
	usecase Usecase
}

func NewHandler(usecase Usecase) *Handler {
	return &Handler{usecase: usecase}
}

type Usecase interface {
	GenerateAltURL(stream io.Reader, size int64, siteURL string) error
}
