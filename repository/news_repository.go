package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/project/go-microservices/domain"
)

type newsRepository struct {
	database *gorm.DB
}
func NewNewsRepository(db *gorm.DB) domain.NewsRepository{
	db.AutoMigrate(&domain.News{})
	return &newsRepository{database: db,}
}

func (tr *newsRepository) Create(c context.Context, news *domain.News) error {
	err := tr.database.Create(news).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return nil
	  }	  
	return err
}

func (tr *newsRepository) FetchByUserID(c context.Context, id string) ([]domain.News, error) {
	var news []domain.News
	err := tr.database.Find(&news, "user_id = ?", id).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return news, nil
	  }	  
	return news, err
}

func (tr *newsRepository) DeleteByUserID(c context.Context, id string) ([]domain.News, error) {
	var user []domain.News
	err := tr.database.Where("user_id = ?", id).Delete(domain.News{}).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return user, nil
	  }	  
	return user, err
}
