package storage

import (
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_service/storage/postgres"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Product() repo.ProductRepoI
	Category() repo.CategoryRepoI
}

type storagePg struct {
	db       *sqlx.DB
	product  repo.ProductRepoI
	category repo.CategoryRepoI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		db:       db,
		product:  postgres.NewProductRepo(db),
		category: postgres.NewCategoryRepo(db),
	}
}

func (s *storagePg) Product() repo.ProductRepoI {
	return s.product
}
func (s *storagePg) Category() repo.CategoryRepoI {
	return s.category
}
