package postgres

import (
	"github.com/sedeotech/backend-interview/internal/config"
	"github.com/sedeotech/backend-interview/internal/storage"
)

type repository struct {
	//db     *bun.DB
	//Logger *zerolog.Logger
}

func NewRepository(cfg config.Config) (storage.Storage, error) {

	return &repository{}, nil
}
