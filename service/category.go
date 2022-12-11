package service

import (
	"context"

	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/catalog_service/genproto/category_service"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/catalog_service/pkg/logger"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/catalog_service/storage"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type categoryService struct {
	log logger.LoggerI
	stg storage.StorageI
	category_service.UnimplementedCategoryServiceServer
}

// NewCategoryService ...
func NewCategoryService(log logger.LoggerI, db *sqlx.DB) *categoryService {
	return &categoryService{
		stg: storage.NewStoragePg(db),
		log: log,
	}
}

// CreateCategory ...
func (s *categoryService) CreateCategory(ctx context.Context, req *category_service.CreateCategoryRequest) (*category_service.Category, error) {
	s.log.Info("---CreateCategory--->", logger.Any("req", req))

	id := uuid.New()

	err := s.stg.Category().CreateCategory(id.String(), req)
	if err != nil {
		s.log.Error("!!!CreateCategory--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	category, err := s.stg.Category().GetCategoryByID(id.String())
	if err != nil {
		s.log.Error("!!!CreateCategory--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &category_service.Category{
		Id:        category.Id,
		Title:     category.Title,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}, nil
}

// GetArticleList ....
func (s *categoryService) GetCategoryList(ctx context.Context, req *category_service.GetCategoryListRequest) (*category_service.GetCategoryListResponse, error) {
	s.log.Info("---GetCategoryList--->", logger.Any("req", req))

	res, err := s.stg.Category().GetCategoryList(int(req.Offset), int(req.Limit), req.Search)
	if err != nil {
		s.log.Error("!!!GetCategoryList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())	}

	return res, nil
}

// GetCategoryByID ....
func (s *categoryService) GetCategoryByID(ctx context.Context, req *category_service.GetCategoryByIdRequest) (*category_service.Category, error) {
	s.log.Info("---GetCategoryByID--->", logger.Any("req", req))
	category, err := s.stg.Category().GetCategoryByID(req.Id)
	if err != nil {
		s.log.Error("!!!GetCategoryByID--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())	
	}

	return category, nil
}

// UpdateCategory ....
func (s *categoryService) UpdateCategory(ctx context.Context, req *category_service.UpdateCategoryRequest) (*category_service.Category, error) {
	s.log.Info("---UpdateCategory--->", logger.Any("req", req))
	err := s.stg.Category().UpdateCategory(req)
	if err != nil {
		s.log.Error("!!!UpdateCategory--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())	
	}

	category, err := s.stg.Category().GetCategoryByID(req.Id)
	if err != nil {
		s.log.Error("!!!UpdateCategory--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())	
	}

	return &category_service.Category{
		Id:        category.Id,
		Title:     category.Title,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}, nil
}

// DeleteCategory ....
func (s *categoryService) DeleteCategory(ctx context.Context, req *category_service.DeleteCategoryRequest) (*category_service.DeleteCategoryResponse, error) {
	s.log.Info("---DeleteCategory--->", logger.Any("req", req))

	category, err := s.stg.Category().GetCategoryByID(req.Id)
	if err != nil {
		s.log.Error("!!!DeleteCategory--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())		}

	err = s.stg.Category().DeleteCategory(category.Id)
	if err != nil {
		s.log.Error("!!!DeleteCategory--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())		
	}

	return &category_service.DeleteCategoryResponse{}, nil
}
