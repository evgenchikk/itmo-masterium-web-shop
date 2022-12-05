package service

import (
	"github.com/evgenchikk/itmo-web-shop/models"
	"github.com/evgenchikk/itmo-web-shop/pkg/repository"
)

type ItemGroupService struct {
	repo repository.ItemGroup
}

func NewItemGroupService(repo repository.ItemGroup) *ItemGroupService {
	return &ItemGroupService{
		repo: repo,
	}
}

func (s *ItemGroupService) GetAllItemGroups() ([]models.ItemGroup, error) {
	return s.repo.GetAll()
}

func (s *ItemGroupService) GetItemGroupById(id int) (models.ItemGroup, error) {
	return s.repo.GetById(id)
}

func (s *ItemGroupService) GetItemGroupsByCategoryId(id int) ([]models.ItemGroup, error) {
	return s.repo.GetByCatgoryId(id)
}
