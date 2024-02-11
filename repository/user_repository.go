package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/project/go-microservices/domain"
)

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	db.AutoMigrate(&domain.User{})
	return &userRepository{
		database: db,
	}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	ur.database.Create(user)
	return nil
}
func (ur *userRepository) Fetch(c context.Context) ([]domain.User, error) {
	var user []domain.User
	ur.database.Create(user)
	return user, nil
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	var user domain.User
	err := ur.database.First(&user, "email = ?", email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, nil
	  }
	  
	return user, err
}

func (ur *userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	var user domain.User
	ur.database.First(&user, "id = ?", id)
	return user, nil
}
