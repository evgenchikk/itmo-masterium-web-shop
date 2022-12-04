package repository

import (
	"context"

	"github.com/evgenchikk/itmo-web-shop/models"

	"github.com/jackc/pgx/v4"
)

type ItemGroupPostgres struct {
	db *pgx.Conn
}

func NewItemGroupPostgres(db *pgx.Conn) *ItemGroupPostgres {
	return &ItemGroupPostgres{
		db: db,
	}
}

func (r *ItemGroupPostgres) GetAll() ([]models.ItemGroup, error) {
	itemGroups := make([]models.ItemGroup, 0)

	query := `SELECT id, category_id, name, image, description FROM item_groups ORDER BY id`
	var itemGroup models.ItemGroup
	if _, err := r.db.QueryFunc(
		context.Background(),
		query,
		[]interface{}{},
		[]interface{}{&itemGroup.Id, &itemGroup.Category_id, &itemGroup.Name, &itemGroup.Image, &itemGroup.Description},
		func(pgx.QueryFuncRow) error {
			itemGroups = append(itemGroups, itemGroup)
			return nil
		},
	); err != nil {
		return nil, err
	}

	return itemGroups, nil
}

func (r *ItemGroupPostgres) GetById(id int) (models.ItemGroup, error) {
	var itemGroup models.ItemGroup

	query := `SELECT id, category_id, name, image, description FROM ITEM_GROUPS WHERE id = $1`
	if err := r.db.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(&itemGroup.Id, &itemGroup.Category_id, &itemGroup.Name, &itemGroup.Image, &itemGroup.Description); err != nil {
		return models.ItemGroup{}, err
	}

	return itemGroup, nil
}
