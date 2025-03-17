package repository

import (
	"context"

	"github.com/YusronAbi/ecomerce-api/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) models.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAllUser(ctx context.Context) ([]*models.User, error) {
	var user []*models.User
	err := r.db.Model(&models.User{}).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (r *UserRepository) GetUserByID(ctx context.Context, id int64) (*models.User, error) {
	user := &models.User{}
	if res := r.db.Model(user).Where("id = ?", id).First(user); res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (r *UserRepository) UpdateUserByID(ctx context.Context, id int64, data map[string]interface{}) (*models.User, error) {
	user := &models.User{}
	res := r.db.Model(&user).Where("id = ?", id).Updates(data)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (r *UserRepository) DeleteUserByID(ctx context.Context, id int64) error {
	user := &models.User{}
	res := r.db.Model(&user).Where("id = ?", id).Delete(user)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
