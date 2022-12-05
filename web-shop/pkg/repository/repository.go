package repository

import (
	"github.com/evgenchikk/itmo-web-shop/models"

	"github.com/jackc/pgx/v4"
)

type Authorization interface {
	// GetUserById() ([]models.User, error)
	SignUp(models.User) (int, error)
}

type ItemGroup interface {
	GetAll() ([]models.ItemGroup, error)
	GetById(id int) (models.ItemGroup, error)
	GetByCatgoryId(id int) ([]models.ItemGroup, error)
}

type Item interface {
	GetAll() ([]models.Item, error)
	GetById(id int) (models.Item, error)
	GetAllByItemGroupId(id int) ([]models.Item, error)
}

type Category interface {
	GetAll() ([]models.Category, error)
	GetById(id int) (models.Category, error)
}

type Repository struct {
	Authorization
	ItemGroup
	Item
	Category
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		ItemGroup:     NewItemGroupPostgres(db),
		Item:          NewItemPostgres(db),
		Category:      NewCategoryPostgres(db),
	}
}
