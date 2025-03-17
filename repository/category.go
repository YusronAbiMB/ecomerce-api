package repository

import (
	"context"

	"github.com/YusronAbi/ecomerce-api/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) models.CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAllCategory(ctx context.Context) ([]*models.Category, error) {
	var categories []*models.Category
	err := r.db.Model(&models.Category{}).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil

}

func (r *CategoryRepository) CreateCategory(ctx context.Context, category *models.Category) (*models.Category, error) {
	if err := r.db.Create(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (r *CategoryRepository) GetCategoryByID(ctx context.Context, id int64) (*models.Category, error) {
	category := &models.Category{}
	if res := r.db.Model(category).Where("id = ?", id).First(category); res.Error != nil {
		return nil, res.Error
	}

	return category, nil
}

func (r *CategoryRepository) UpdateCategoryByID(ctx context.Context, id int64, data map[string]interface{}) (*models.Category, error) {
	category := &models.Category{}
	res := r.db.Model(&category).Where("id = ?", id).Updates(data)

	if res.Error != nil {
		return nil, res.Error
	}

	return category, nil
}

func (r *CategoryRepository) DeleteCategoryByID(ctx context.Context, id int64) error {
	category := &models.Category{}
	res := r.db.Model(&category).Where("id = ?", id).Delete(category)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
