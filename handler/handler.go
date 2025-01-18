package handler

type Handler struct {
	usecase Usecase
}

func NewHandler(usecase Usecase) *Handler {
	return &Handler{usecase: usecase}
}

type Usecase interface {
	GenerateAltURL(url string, ogp []byte) error
	GetRedirectPage() error
}
