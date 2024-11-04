package service

import (
	"context"
	"flavor/internal/adapters/db/mongo/storage"
	"flavor/internal/domain/entity"
)

type CategoryService struct {
	categoryStorage storage.CategoryStorager
}

func NewCategoryService(categoryStorage storage.CategoryStorager) *CategoryService {
	return &CategoryService{
		categoryStorage: categoryStorage,
	}
}

func (categoryService CategoryService) GetCategoryByPath(ctx context.Context, path string) (entity.Category, error) {
	return categoryService.categoryStorage.GetByPath(ctx, path)
}

func (categoryService CategoryService) GetCategoriesAllMain(ctx context.Context) ([]entity.Category, error) {
	return categoryService.categoryStorage.GetAllMain(ctx)
}

func (categoryService CategoryService) AddCategory(ctx context.Context, category entity.Category) error {
	return categoryService.categoryStorage.Create(ctx, category)
}
