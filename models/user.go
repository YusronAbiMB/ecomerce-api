package models

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role" gorm:"default:user"`
}

func (u *User) AfterCreate(db *gorm.DB) (err error) {
	if u.ID == 1 {
		db.Model(u).Update("role", "admin")
	}
	return
}

type UserRepository interface {
	GetAllUser(ctx context.Context) ([]*User, error)
	GetUserByID(ctx context.Context, id int64) (*User, error)
	UpdateUserByID(ctx context.Context, id int64, data map[string]interface{}) (*User, error)
	DeleteUserByID(ctx context.Context, id int64) error
}
