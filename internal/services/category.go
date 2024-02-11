package services

import (
	"context"

	"github.com/almirpask/go_grpc/internal/database"
	"github.com/almirpask/go_grpc/internal/pb"
)

type CateogryService struct {
	pb.UnimplementedCateogryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CateogryService {
	return &CateogryService{CategoryDB: categoryDB}
}

func (c *CateogryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}
	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{Category: categoryResponse}, nil
}
