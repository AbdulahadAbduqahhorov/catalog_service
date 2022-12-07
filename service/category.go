package service

import (
	"context"

	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_service/genproto/category_service"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_service/storage"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


type categoryService struct {
	stg storage.StorageI
	category_service.UnimplementedCategoryServiceServer
}

// NewCategoryService ...
func NewCategoryService(stg storage.StorageI) *categoryService {
	return &categoryService{
		stg: stg,
	}
}

// CreateCategory ...
func (s *categoryService) CreateCategory(ctx context.Context, req *category_service.CreateCategoryRequest) (*category_service.Category, error) {
	id := uuid.New()

	err := s.stg.Category().CreateCategory(id.String(), req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.CreateCategory: %s", err.Error())
	}

	category, err := s.stg.Category().GetCategoryByID(id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.Stg.GetCategoryByID: %s", err.Error())
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
	res, err := s.stg.Category().GetCategoryList(int(req.Offset), int(req.Limit), req.Search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteCategory: %s", err.Error())
	}

	return res, nil
}

// GetCategoryByID ....
func (s *categoryService) GetCategoryByID(ctx context.Context, req *category_service.GetCategoryByIdRequest) (*category_service.Category, error) {
	category, err := s.stg.Category().GetCategoryByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetCategoryByID: %s", err.Error())
	}

	return category, nil
}

// UpdateCategory ....
func (s *categoryService) UpdateCategory(ctx context.Context, req *category_service.UpdateCategoryRequest) (*category_service.Category, error) {
	err := s.stg.Category().UpdateCategory(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.UpdateCategory: %s", err.Error())
	}

	category, err := s.stg.Category().GetCategoryByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetCategoryByID: %s", err.Error())
	}

	return &category_service.Category{
		Id:        category.Id,
		Title:     category.Title,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}, nil
}

// DeleteCategory ....
func (s *categoryService) DeleteCategory(ctx context.Context, req *category_service.DeleteCategoryRequest) (*category_service.Category, error) {
	category, err := s.stg.Category().GetCategoryByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetCategoryByID: %s", err.Error())
	}

	err = s.stg.Category().DeleteCategory(category.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteCategory: %s", err.Error())
	}

	return &category_service.Category{
		Id:        category.Id,
		Title:     category.Title,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}, nil
}
