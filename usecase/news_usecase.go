package usecase

import (
	"context"
	"time"

	"backend-task/domain"
)

type newsUsecase struct {
	newsRepository domain.NewsRepository
	contextTimeout time.Duration
}

func NewNewsUsecase(newsRepository domain.NewsRepository, timeout time.Duration) domain.NewsUsecase {
	return &newsUsecase{
		newsRepository: newsRepository,
		contextTimeout: timeout,
	}
}

func (tu *newsUsecase) Create(c context.Context, news *domain.News) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.newsRepository.Create(ctx, news)
}

func (tu *newsUsecase) FetchByUserID(c context.Context, userID string) ([]domain.News, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.newsRepository.FetchByUserID(ctx, userID)
}

func (tu *newsUsecase) DeleteByUserID(c context.Context, userID string) ([]domain.News, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.newsRepository.FetchByUserID(ctx, userID)
}
