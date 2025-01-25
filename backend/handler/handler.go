package handler

import (
	"io"

	"github.com/ponyo877/ogper/domain"
)

type Handler struct {
	usecase Usecase
}

func NewHandler(usecase Usecase) *Handler {
	return &Handler{usecase: usecase}
}

type Usecase interface {
	GenerateOGPPage(title, description, name, siteURL, userHash string, file io.Reader, size int64) (string, error)
	GetOGPPage(hash string) (string, error)
	ListSitesByUserID(userHash string) ([]*domain.Site, error)
}
