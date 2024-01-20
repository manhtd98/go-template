package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionNew = "news"
)

type News struct {
	ID     primitive.ObjectID `bson:"_id" json:"-"`
	Title  string             `bson:"title" form:"title" binding:"required" json:"title"`
	UserID primitive.ObjectID `bson:"userID" json:"-"`
	CreatedAt primitive.DateTime `bson:"createdAt" json:"-"`
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
