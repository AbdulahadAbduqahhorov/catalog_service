package service

import (
	"context"

	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/catalog_service/genproto/product_service"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/catalog_service/storage"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type productService struct {
	product_service.UnimplementedProductServiceServer
	stg storage.StorageI
}

func NewProductService(db *sqlx.DB) *productService {
	return &productService{
		stg: storage.NewStoragePg(db),
	}
}

func (s *productService) CreateProduct(ctx context.Context, req *product_service.CreateProductRequest) (*product_service.CreateProductResponse, error) {

	// _, err := s.stg.Category().GetCategoryById(req.CategoryId)
	// if err != nil {
	// 	return nil, status.Errorf(codes.NotFound, "method GetCategoryById: %v",err)

	// }
	id := uuid.New().String()

	err := s.stg.Product().CreateProduct(id, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "method CreateProduct: %v", err)
	}

	_, err = s.stg.Product().GetProductById(id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "method GetProductById: %v", err)

	}

	return &product_service.CreateProductResponse{
		Id: id,
	}, nil
}
func (s *productService) GetProductList(ctx context.Context, req *product_service.GetProductListRequest) (*product_service.GetProductListResponse, error) {

	res, err := s.stg.Product().GetProductList(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "method GetProductList: %v", err)

	}
	return res, nil
}
func (s *productService) GetProductById(ctx context.Context, req *product_service.GetProductByIdRequest) (*product_service.GetProductByIdResponse, error) {

	res, err := s.stg.Product().GetProductById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "method GetProductById: %v", err)

	}
	return res, nil

}
func (s *productService) UpdateProduct(ctx context.Context, req *product_service.UpdateProductRequest) (*product_service.Product, error) {

	rowsAffected, err := s.stg.Product().UpdateProduct(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "method UpdateProduct: %v", err)
	}
	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}
	product, err := s.stg.Product().GetProductById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "method GetProductById: %v", err)

	}
	return &product_service.Product{
		Id:         product.Id,
		Title:      product.Title,
		Desc:       product.Desc,
		Quantity:   product.Quantity,
		Price:      product.Price,
		CategoryId: product.Category.Id,
		CreatedAt:  product.CreatedAt,
		UpdatedAt:  product.UpdatedAt,
	}, nil
}
func (s *productService) DeleteProduct(ctx context.Context, req *product_service.DeleteProductRequest) (*product_service.Empty, error) {

	rowsAffected, err := s.stg.Product().DeleteProduct(req.Id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "method DeleteProduct : %v", err)

	}
	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return &product_service.Empty{}, nil

}
