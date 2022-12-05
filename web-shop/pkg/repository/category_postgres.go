package repository

import (
	"context"

	"github.com/evgenchikk/itmo-web-shop/models"
	"github.com/jackc/pgx/v4"
)

type CategoryPostgres struct {
	db *pgx.Conn
}

func NewCategoryPostgres(db *pgx.Conn) *CategoryPostgres {
	return &CategoryPostgres{
		db: db,
	}
}

func (r *CategoryPostgres) GetAll() ([]models.Category, error) {
	categories := make([]models.Category, 0)

	query := `SELECT id, name FROM item_categories ORDER BY id`
	var category models.Category
	if _, err := r.db.QueryFunc(
		context.Background(),
		query,
		[]interface{}{},
		[]interface{}{&category.Id, &category.Name},
		func(pgx.QueryFuncRow) error {
			categories = append(categories, category)
			return nil
		},
	); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryPostgres) GetById(id int) (models.Category, error) {
	var category models.Category

	query := `SELECT id, name FROM item_categories WHERE id = $1`
	if err := r.db.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(&category.Id, &category.Name); err != nil {
		return models.Category{}, err
	}

	return category, nil
}
