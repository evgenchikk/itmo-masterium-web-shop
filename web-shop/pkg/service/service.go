package service

import (
	"github.com/evgenchikk/itmo-web-shop/models"
	"github.com/evgenchikk/itmo-web-shop/pkg/repository"
)

type Authorization interface {
	SignUp(models.User) (int, error)
	// SignIn(models.User) (int, error)
	// GenerateToken() (int, error)
	// ParseToken() (int, error)
}

type ItemGroup interface {
	GetAllItemGroups() ([]models.ItemGroup, error)
	GetItemGroupById(id int) (models.ItemGroup, error)
	GetItemGroupsByCategoryId(id int) ([]models.ItemGroup, error)
}

type Item interface {
	GetAllItems() ([]models.Item, error)
	GetItemById(id int) (models.Item, error)
	GetAllItemsByItemGroupId(id int) ([]models.Item, error)
}

type Category interface {
	GetAllCategories() ([]models.Category, error)
	GetCategoryById(id int) (models.Category, error)
}

type Service struct {
	Authorization
	ItemGroup
	Item
	Category
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		ItemGroup:     NewItemGroupService(repo.ItemGroup),
		Item:          NewItemService(repo.Item),
		Category:      NewCategoryService(repo.Category),
	}
}
