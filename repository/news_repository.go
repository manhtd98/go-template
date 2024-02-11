package repository

import (
	"context"

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
	tr.database.Create(news)
	return nil
}

func (tr *newsRepository) FetchByUserID(c context.Context, id string) ([]domain.News, error) {
	var user []domain.News
	return user, nil
}

func (tr *newsRepository) DeleteByUserID(c context.Context, id string) ([]domain.News, error) {
	var user []domain.News
	return user, nil
}
