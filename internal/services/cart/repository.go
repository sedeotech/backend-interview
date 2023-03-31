package cart

import "github.com/sedeotech/backend-interview/internal/storage"

type Repository interface {
	storage.CartStore
}
