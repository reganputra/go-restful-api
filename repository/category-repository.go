package repository

import (
	"context"
	"database/sql"
	"go-restful-api/model/entity"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category
	Update(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category
	Delete(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (entity.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Category
}
