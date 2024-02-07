package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/project/go-microservices/domain"
	"github.com/project/go-microservices/db"
)

type newsRepository struct {
	database   db.Database
	collection string
}

func NewNewsRepository(db db.Database, collection string) domain.NewsRepository {
	return &newsRepository{
		database:   db,
		collection: collection,
	}
}

func (tr *newsRepository) Create(c context.Context, news *domain.News) error {
	collection := tr.database.Collection(tr.collection)
	news.CreatedAt = primitive.NewDateTimeFromTime(time.Now().UTC())
	_, err := collection.InsertOne(c, news)

	return err
}

func (tr *newsRepository) FetchByUserID(c context.Context, userID string) ([]domain.News, error) {
	collection := tr.database.Collection(tr.collection)

	var news []domain.News

	idHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return news, err
	}

	cursor, err := collection.Find(c, bson.M{"userID": idHex})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &news)
	if news == nil {
		return []domain.News{}, err
	}

	return news, err
}

func (tr *newsRepository) DeleteByUserID(c context.Context, userID string) ([]domain.News, error) {
	collection := tr.database.Collection(tr.collection)

	var news []domain.News

	idHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return news, err
	}

	cursor, err := collection.Find(c, bson.M{"userID": idHex})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &news)
	if news == nil {
		return []domain.News{}, err
	}

	return news, err
}
