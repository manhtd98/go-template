package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	UserID  string
	CreatedAt time.Time
	Title string
	Content string
}

type NewsRepository interface {
	Create(c context.Context, news *News) error
	FetchByUserID(c context.Context, userID string) ([]News, error)
	DeleteByUserID(c context.Context, userID string) ([]News, error)
}

type NewsUsecase interface {
	Create(c context.Context, news *News) error
	FetchByUserID(c context.Context, userID string) ([]News, error)
	DeleteByUserID(c context.Context, userID string) ([]News, error)
}
