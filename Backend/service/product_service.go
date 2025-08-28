package service

import (
	"inventory_backend/model"
	"inventory_backend/repository"
)

type ProductService interface {
	Create(product model.Product) (model.Product, error)
	FindAll(status string) ([]model.Product, error)
	FindByID(id int) (*model.Product, error)
	Update(product *model.Product) (model.Product, error)
	Delete(id int) (model.Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *productService {
	return &productService{repo: repo}
}

func (s *productService) Create(product model.Product) (model.Product, error) {
	return s.repo.Create(product)
}
func (s *productService) FindAll(status string) ([]model.Product, error) {
	return s.repo.FindAll(status)
}
func (s *productService) FindByID(id int) (*model.Product, error) {
	return s.repo.FindByID(id)
}
func (s *productService) Update(product *model.Product) (model.Product, error) {
	return s.repo.Update(product)
}
func (s *productService) Delete(id int) (model.Product, error) {
	return s.repo.Delete(id)
}
