package service

import (
	"github.com/evgenchikk/itmo-web-shop/models"
	"github.com/evgenchikk/itmo-web-shop/pkg/repository"
)

type ItemService struct {
	repo repository.Item
}

func NewItemService(repo repository.Item) *ItemService {
	return &ItemService{
		repo: repo,
	}
}

func (s *ItemService) GetAllItems() ([]models.Item, error) {
	return s.repo.GetAll()
}

func (s *ItemService) GetItemById(id int) (models.Item, error) {
	return s.repo.GetById(id)
}

func (s *ItemService) GetAllItemsByItemGroupId(id int) ([]models.Item, error) {
	return s.repo.GetAllByItemGroupId(id)
}
