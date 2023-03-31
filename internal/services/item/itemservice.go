package item

type ItemService interface {
}

type service struct {
	repository Repository
}

func NewItemService(r Repository) ItemService {

	return &service{
		repository: r,
	}
}
