package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"go-restful-api/exception"
	"go-restful-api/helper"
	"go-restful-api/model/entity"
	"go-restful-api/model/web"
	"go-restful-api/repository"
)

type CategoryServiceImplementation struct {
	Repository repository.CategoryRepository
	Db         *sql.DB
	Validate   *validator.Validate
}

func NewCategoryService(repository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImplementation{
		Repository: repository,
		Db:         db,
		Validate:   validate,
	}
}

// Every method here is Transactional

func (service *CategoryServiceImplementation) Create(ctx context.Context, req web.CategoryCreateRequest) web.CategoryResponse {

	err := service.Validate.Struct(req) // Validate the request
	helper.PanicIfError(err)

	tx, err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := entity.Category{
		Name: req.Name,
	}

	category = service.Repository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImplementation) Update(ctx context.Context, req web.CategoryUpdateRequest) web.CategoryResponse {

	err := service.Validate.Struct(req) // Validate the request
	helper.PanicIfError(err)

	tx, err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.Repository.FindById(ctx, tx, req.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = req.Name

	category = service.Repository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImplementation) Delete(ctx context.Context, categoryId int) {
	tx, err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.Repository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.Repository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImplementation) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.Repository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImplementation) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.Repository.FindAll(ctx, tx)

	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, helper.ToCategoryResponse(category))
	}
	return categoryResponses
}
