package models

import (
	"context"
)

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CategoryID  int64   `json:"category_id"`
}

type ProductRepository interface {
	GetAllProduct(ctx context.Context) ([]*Product, error)
	GetProductByID(ctx context.Context, id int64) (*Product, error)
	CreateProduct(ctx context.Context, category *Product) (*Product, error)
	UpdateProductByID(ctx context.Context, id int64, data map[string]interface{}) (*Product, error)
	DeleteProductByID(ctx context.Context, id int64) error
}
