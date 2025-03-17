package repository

import (
	"context"

	"github.com/YusronAbi/ecomerce-api/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) models.ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAllProduct(ctx context.Context) ([]*models.Product, error) {
	var products []*models.Product
	err := r.db.Model(&models.Product{}).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil

}

func (r *ProductRepository) CreateProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	if err := r.db.Create(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id int64) (*models.Product, error) {
	product := &models.Product{}
	if res := r.db.Model(product).Where("id = ?", id).First(product); res.Error != nil {
		return nil, res.Error
	}

	return product, nil
}

func (r *ProductRepository) UpdateProductByID(ctx context.Context, id int64, data map[string]interface{}) (*models.Product, error) {
	product := &models.Product{}
	res := r.db.Model(&product).Where("id = ?", id).Updates(data)

	if res.Error != nil {
		return nil, res.Error
	}

	return product, nil
}

func (r *ProductRepository) DeleteProductByID(ctx context.Context, id int64) error {
	product := &models.Product{}
	res := r.db.Model(&product).Where("id = ?", id).Delete(product)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
