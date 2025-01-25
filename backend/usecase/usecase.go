package usecase

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ponyo877/ogper/domain"
)

type Usecase struct {
	repository Repository
}

var fileDomain = os.Getenv("FILE_DOMAIN")
var ogpPageDomain = os.Getenv("OGP_PAGE_DOMAIN")

type Repository interface {
	PutFile(file []byte, filename, contentType string) error
	CreateSite(site *domain.Site) error
	GetSite(hash string) (*domain.Site, error)
	GetHtml(site *domain.Site) (string, error)
}

func NewUsecase(repository Repository) *Usecase {
	return &Usecase{repository: repository}
}

func (u *Usecase) GenerateOGPPage(title, description, name, siteURL string, file io.Reader, size int64) (string, error) {
	filedata := make([]byte, size)
	if _, err := file.Read(filedata); err != nil {
		return "", err
	}
	hash := domain.NewHash().String()
	contentType := http.DetectContentType(filedata)
	var contentTypeToExtension = map[string]string{
		"image/webp": "webp",
		"image/png":  "png",
		"image/jpeg": "jpg",
	}
	ext, ok := contentTypeToExtension[contentType]
	if !ok {
		return "", fmt.Errorf("unsupported content type")
	}
	filename := hash + "." + ext

	if err := u.repository.PutFile(filedata, filename, contentType); err != nil {
		return "", err
	}
	imageURL := fileDomain + "/" + filename
	site := domain.NewSite(hash, title, description, name, siteURL, imageURL)
	if err := u.repository.CreateSite(site); err != nil {
		return "", err
	}
	return ogpPageDomain + "/" + hash, nil
}

func (u *Usecase) GetOGPPage(hash string) (string, error) {
	site, err := u.repository.GetSite(hash)
	if err != nil {
		return "", err
	}
	html, err := u.repository.GetHtml(site)
	if err != nil {
		return "", err
	}
	return html, nil
}
