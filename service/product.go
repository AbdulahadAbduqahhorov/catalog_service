package service

import (
	"context"

	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/catalog_service/genproto/product_service"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/catalog_service/pkg/logger"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/catalog_service/storage"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type productService struct {
	log logger.LoggerI
	stg storage.StorageI
	product_service.UnimplementedProductServiceServer
}

func NewProductService(log logger.LoggerI, db *sqlx.DB) *productService {
	return &productService{
		log: log,
		stg: storage.NewStoragePg(db),
	}
}

func (s *productService) CreateProduct(ctx context.Context, req *product_service.CreateProductRequest) (*product_service.CreateProductResponse, error) {
	s.log.Info("---CreateProduct--->", logger.Any("req", req))
	// _, err := s.stg.Category().GetCategoryById(req.CategoryId)
	// if err != nil {
	// 	return nil, status.Errorf(codes.NotFound, "method GetCategoryById: %v",err)

	// }
	id := uuid.New().String()

	err := s.stg.Product().CreateProduct(id, req)
	if err != nil {
		s.log.Error("!!!CreateProduct--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())	
	}

	_, err = s.stg.Product().GetProductById(id)
	if err != nil {
		s.log.Error("!!!GetProductById--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())	
	}

	return &product_service.CreateProductResponse{
		Id: id,
	}, nil	
}
func (s *productService) GetProductList(ctx context.Context, req *product_service.GetProductListRequest) (*product_service.GetProductListResponse, error) {
	s.log.Info("---GetProductList--->", logger.Any("req", req))
	res, err := s.stg.Product().GetProductList(req)
	if err != nil {
		s.log.Error("!!!GetProductList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return res, nil
}
func (s *productService) GetProductById(ctx context.Context, req *product_service.GetProductByIdRequest) (*product_service.GetProductByIdResponse, error) {
	s.log.Info("---GetProductById--->", logger.Any("req", req))

	res, err := s.stg.Product().GetProductById(req.Id)
	if err != nil {
		s.log.Error("!!!GetProductById--->", logger.Error(err))

		return nil, status.Error(codes.NotFound, err.Error())
	}
	return res, nil

}
func (s *productService) UpdateProduct(ctx context.Context, req *product_service.UpdateProductRequest) (*product_service.Product, error) {
	s.log.Info("---UpdateProduct--->", logger.Any("req", req))

	rowsAffected, err := s.stg.Product().UpdateProduct(req)
	if err != nil {
		s.log.Error("!!!UpdateProduct--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}
	product, err := s.stg.Product().GetProductById(req.Id)
	if err != nil {
		s.log.Error("!!!UpdateProduct--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
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
	s.log.Info("---DeleteProduct--->", logger.Any("req", req))

	rowsAffected, err := s.stg.Product().DeleteProduct(req.Id)

	if err != nil {
		s.log.Error("!!!DeleteProduct--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return &product_service.Empty{}, nil

}
