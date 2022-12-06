package service

import (
	"context"

	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_service/genproto/category_service"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_service/storage"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// StorageI ...
type StorageI interface {
	CreateCategory(id string, entity *category_service.CreateCategoryRequest) error
	GetCategoryByID(id string) (*category_service.GetCategoryIDResponse, error)
	GetCategoryList(offset, limit int, search string) (resp *category_service.GetCategoryListResponse, err error)
	UpdateCategory(entity *category_service.UpdateCategoryRequest) error
	DeleteCategory(id string) error
}

// categoryService := category.NewCategoryService(stg)
// category_service.RegisterCategoryServiceServer(srv, categoryService)

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

	err := s.stg.CreateCategory(id.String(), req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.CreateCategory: %s", err.Error())
	}

	category, err := s.stg.GetCategoryByID(id.String())
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
func (s *categoryService) GetCategoryList(ctx context.Context, req *category_service.GetCategoryRequest) (*category_service.GetCategoryListResponse, error) {
	res, err := s.stg.GetCategoryList(int(req.Offset), int(req.Limit), req.Search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteCategory: %s", err.Error())
	}

	return res, nil
}

// GetCategoryByID ....
func (s *categoryService) GetCategoryByID(ctx context.Context, req *category_service.GetCategoryByIDRequest) (*category_service.GetCategoryByIDResponse, error) {
	category, err := s.stg.GetCategoryByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetCategoryByID: %s", err.Error())
	}

	return category, nil
}

// UpdateCategory ....
func (s *categoryService) UpdateCategory(ctx context.Context, req *category_service.UpdateCategoryRequest) (*category_service.Category, error) {
	err := s.stg.UpdateCategory(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.UpdateCategory: %s", err.Error())
	}

	category, err := s.stg.GetCategoryByID(req.Id)
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
	category, err := s.stg.GetCategoryByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetCategoryByID: %s", err.Error())
	}

	err = s.stg.DeleteCategory(category.Id)
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
