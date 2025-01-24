package domain

type Site struct {
	hash        string
	title       string
	description string
	name        string
	siteURL     string
	imageURL    string
}

func NewSite(hash, title, description, name, siteURL, imageURL string) *Site {
	return &Site{
		hash:        hash,
		title:       title,
		description: description,
		name:        name,
		siteURL:     siteURL,
		imageURL:    imageURL,
	}
}

func (s *Site) Hash() string {
	return s.hash
}

func (s *Site) Title() string {
	return s.title
}

func (s *Site) Description() string {
	return s.description
}

func (s *Site) Name() string {
	return s.name
}

func (s *Site) SiteURL() string {
	return s.siteURL
}

func (s *Site) ImageURL() string {
	return s.imageURL
}
