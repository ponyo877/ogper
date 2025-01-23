package usecase

import (
	"io"
	"net/http"

	"github.com/oklog/ulid/v2"
	"github.com/ponyo877/ogper/domain"
)

type Usecase struct {
	repository Repository
}

type Repository interface {
	PutFile(file []byte, filename, contentType string) error
	CreateSite(site *domain.Site) error
	GetSite(hash string) (*domain.Site, error)
	GetHtml(site *domain.Site) (string, error)
}

func NewUsecase(repository Repository) *Usecase {
	return &Usecase{repository: repository}
}

func (u *Usecase) GenerateOGPPage(title, description, name, siteURL string, file io.Reader, size int64) error {
	filedata := make([]byte, size)
	if _, err := file.Read(filedata); err != nil {
		return err
	}
	id := ulid.MustNew(ulid.Now(), nil).String()
	contentType := http.DetectContentType(filedata)
	var contentTypeToExtension = map[string]string{
		"image/webp": "webp",
		"image/png":  "png",
		"image/jpeg": "jpg",
	}
	ext := contentTypeToExtension[contentType]
	filename := id + "." + ext

	if err := u.repository.PutFile(filedata, filename, contentType); err != nil {
		return err
	}
	imageURL := "https://r2.folks-chat.com/" + filename
	hash := domain.NewHash().String()
	site := domain.NewSite(hash, title, description, name, siteURL, imageURL)
	return u.repository.CreateSite(site)
}

func (u *Usecase) GetOGPPage(hash string) (string, string, error) {
	site, err := u.repository.GetSite(hash)
	if err != nil {
		return "", "", err
	}
	html, err := u.repository.GetHtml(site)
	if err != nil {
		return "", "", err
	}
	return site.SiteURL(), html, nil
}
