package repository

import (
	"user-management-2/entity"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user entity.User) error
	GetAll() ([]entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) Create(user entity.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u UserRepository) GetAll() ([]entity.User, error) {
	var users []entity.User
	if err := u.db.Debug().Find(&users).Error; err != nil {
		return nil, nil
	}
	return users, nil
}
