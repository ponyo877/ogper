package repository

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetRedirectPage() error {
	return nil
}

func (r *Repository) GenerateAltURL() error {
	return nil
}
