package service

import (
	"context"
	"go-restful-api/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, req web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, req web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
