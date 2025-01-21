package usecase

import (
	"io"
	"net/http"

	"github.com/oklog/ulid/v2"
)

type Usecase struct {
	repository Repository
}

type Repository interface {
	GetRedirectPage() error
	PutFile(file []byte, filename, contentType string) error
}

func NewUsecase(repository Repository) *Usecase {
	return &Usecase{repository: repository}
}

func (u *Usecase) GenerateAltURL(file io.Reader, size int64) error {
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
	return nil
}

func (u *Usecase) GetRedirectPage() error {
	return nil
}
