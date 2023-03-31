package item

import "github.com/sedeotech/backend-interview/internal/storage"

type Repository interface {
	storage.ItemStore
}
