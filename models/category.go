package models

import (
	"context"
)

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type CategoryRepository interface {
	GetAllCategory(ctx context.Context) ([]*Category, error)
	GetCategoryByID(ctx context.Context, id int64) (*Category, error)
	CreateCategory(ctx context.Context, category *Category) (*Category, error)
	UpdateCategoryByID(ctx context.Context, id int64, data map[string]interface{}) (*Category, error)
	DeleteCategoryByID(ctx context.Context, id int64) error
}
