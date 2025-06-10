package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-restful-api/helper"
	"go-restful-api/model/entity"
)

type CategoryRepositoryImplementation struct {
}

func (c *CategoryRepositoryImplementation) Save(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {

	insert := "INSERT INTO category (name) VALUES (?)"
	result, err := tx.ExecContext(ctx, insert, category.Name)
	helper.PanicIfError(err) // Handle error if the query fails

	insertedId, err := result.LastInsertId() // Auto increment ID
	helper.PanicIfError(err)

	category.Id = int(insertedId)
	return category
}

func (c *CategoryRepositoryImplementation) Update(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	update := "UPDATE category SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, update, category.Name, category.Id)
	helper.PanicIfError(err) // Handle error if the query fails

	return category
}

func (c *CategoryRepositoryImplementation) Delete(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	deleteQuery := "DELETE FROM category WHERE id = ?"
	_, err := tx.ExecContext(ctx, deleteQuery, category.Id)
	helper.PanicIfError(err) // Handle error if the query fails
	return category
}

func (c *CategoryRepositoryImplementation) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (entity.Category, error) {
	findByIdQuery := "SELECT id, name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, findByIdQuery, categoryId)
	helper.PanicIfError(err) // Handle error if the query fails

	category := entity.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}

}

func (c *CategoryRepositoryImplementation) FindAll(ctx context.Context, tx *sql.Tx) []entity.Category {
	findAllQuery := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, findAllQuery)
	helper.PanicIfError(err) // Handle error if the query fails

	var categories []entity.Category
	for rows.Next() {
		category := entity.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
