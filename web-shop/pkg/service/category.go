package service

import (
	"github.com/evgenchikk/itmo-web-shop/models"
	"github.com/evgenchikk/itmo-web-shop/pkg/repository"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) GetCategoryById(id int) (models.Category, error) {
	return s.repo.GetById(id)
}
