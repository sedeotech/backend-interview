package storage

type Storage interface {
	CartStore
	ItemStore
}

type CartStore interface{}
type ItemStore interface{}
