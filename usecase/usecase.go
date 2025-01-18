package usecase

type Usecase struct {
	repository Repository
}

type Repository interface {
	GetRedirectPage() error
	GenerateAltURL() error
}

func NewUsecase(repository Repository) *Usecase {
	return &Usecase{repository: repository}
}

func (u *Usecase) GenerateAltURL(url string, ogp []byte) error {
	return nil
}

func (u *Usecase) GetRedirectPage() error {
	return nil
}
