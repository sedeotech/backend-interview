package cart

type CartService interface {
}

type service struct {
	repository Repository
}

func NewCartService(r Repository) CartService {

	return &service{
		repository: r,
	}
}
