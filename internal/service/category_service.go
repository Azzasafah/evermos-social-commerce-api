package service

import (
	"evermos-backend/internal/dto"
	"evermos-backend/internal/model"
	"evermos-backend/internal/repository"
)

type CategoryService interface {
	GetAllCategories() ([]model.Category, error)
	GetCategoryByID(id uint) (*model.Category, error)
	CreateCategory(req *dto.CategoryRequest) error
	UpdateCategory(id uint, req *dto.CategoryRequest) error
	DeleteCategory(id uint) error
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo: categoryRepo}
}

func (s *categoryService) GetAllCategories() ([]model.Category, error) {
	categories, err := s.categoryRepo.FindAll()
	if err != nil {
		return nil, err
	}
	if categories == nil {
		categories = []model.Category{}
	}
	return categories, nil
}

func (s *categoryService) GetCategoryByID(id uint) (*model.Category, error) {
	return s.categoryRepo.FindByID(id)
}

func (s *categoryService) CreateCategory(req *dto.CategoryRequest) error {
	category := model.Category{
		NamaCategory: req.NamaCategory,
	}
	return s.categoryRepo.Create(&category)
}

func (s *categoryService) UpdateCategory(id uint, req *dto.CategoryRequest) error {
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return err
	}

	category.NamaCategory = req.NamaCategory
	return s.categoryRepo.Update(category)
}

func (s *categoryService) DeleteCategory(id uint) error {
	return s.categoryRepo.Delete(id)
}
