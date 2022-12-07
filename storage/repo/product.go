package repo

import "github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/catalog_service/genproto/product_service"

type ProductRepoI interface {
	CreateProduct(id string, req *product_service.CreateProductRequest) error
	GetProductList(req *product_service.GetProductListRequest) (*product_service.GetProductListResponse, error)
	GetProductById(id string) (*product_service.GetProductByIdResponse, error)
	UpdateProduct(req *product_service.UpdateProductRequest) (int64, error)
	DeleteProduct(id string) (int64, error)
}
