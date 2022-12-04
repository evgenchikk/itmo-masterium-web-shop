package repository

import (
	"context"

	"github.com/evgenchikk/itmo-web-shop/models"

	"github.com/jackc/pgx/v4"
)

type ItemPostgres struct {
	db *pgx.Conn
}

func NewItemPostgres(db *pgx.Conn) *ItemPostgres {
	return &ItemPostgres{
		db: db,
	}
}

func (r *ItemPostgres) GetAll() ([]models.Item, error) {
	items := make([]models.Item, 0)

	query := `SELECT id, item_group_id, price, image, feature FROM items ORDER BY id`
	var item models.Item
	if _, err := r.db.QueryFunc(
		context.Background(),
		query,
		[]interface{}{},
		[]interface{}{&item.Id, &item.Item_group_id, &item.Price, &item.Image, &item.Feature},
		func(pgx.QueryFuncRow) error {
			items = append(items, item)
			return nil
		},
	); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *ItemPostgres) GetById(id int) (models.Item, error) {
	var item models.Item

	query := `SELECT id, item_group_id, price, image, feature FROM items WHERE id = $1`
	if err := r.db.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(&item.Id, &item.Item_group_id, &item.Price, &item.Image, &item.Feature); err != nil {
		return models.Item{}, err
	}

	return item, nil
}

func (r *ItemPostgres) GetAllByItemGroupId(id int) ([]models.Item, error) {
	items := make([]models.Item, 0)

	query := `SELECT id, item_group_id, price, image, feature FROM items WHERE item_group_id = $1 ORDER BY id`
	var item models.Item
	if _, err := r.db.QueryFunc(
		context.Background(),
		query,
		[]interface{}{id},
		[]interface{}{&item.Id, &item.Item_group_id, &item.Price, &item.Image, &item.Feature},
		func(pgx.QueryFuncRow) error {
			items = append(items, item)
			return nil
		},
	); err != nil {
		return nil, err
	}

	return items, nil
}
