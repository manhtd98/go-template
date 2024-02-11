package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	UserID  string `bson:"user_id"`
	CreatedAt time.Time `bson:"created_at"`
	Title string `bson:"title"`
	Content string `bson:"content"`
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
