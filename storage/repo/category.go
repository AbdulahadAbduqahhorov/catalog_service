package repo

import "github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/catalog_service/genproto/category_service"

type CategoryRepoI interface {
	CreateCategory(id string, req *category_service.CreateCategoryRequest) error
	GetCategoryList(offset, limit int, search string) (*category_service.GetCategoryListResponse, error)
	GetCategoryById(id string) (*category_service.Category, error)
	UpdateCategory(req *category_service.UpdateCategoryRequest) error
	DeleteCategory(id string) error
}
