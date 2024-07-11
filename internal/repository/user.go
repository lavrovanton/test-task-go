package repository

import (
	"test-task-go/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetAdmin() (model.User, error) {
	var user model.User
	result := r.db.Where("name = ?", "admin").First(&user)
	return user, result.Error
}
