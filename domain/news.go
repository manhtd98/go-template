package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	ID        int32     `json:"id" gorm:"primaryKey;autoIncrement:true"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
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
