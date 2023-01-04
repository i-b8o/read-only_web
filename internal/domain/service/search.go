package service

import (
	"context"
	"read-only_web/internal/domain/entity"
)

type SearchStorage interface {
	Search(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error)
	SearchLike(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error)
	ChSearch(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error)
	RegSearch(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error)
	PSearch(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error)
}

type searchService struct {
	storage SearchStorage
}

func NewSearchService(storage SearchStorage) *searchService {
	return &searchService{storage: storage}
}

func (ss searchService) Search(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error) {
	return ss.storage.Search(ctx, searchQuery, params...)
}

func (ss searchService) SearchLike(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error) {
	return ss.storage.SearchLike(ctx, searchQuery, params...)
}

func (ss searchService) RegSearch(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error) {
	return ss.storage.RegSearch(ctx, searchQuery, params...)
}

func (ss searchService) ChSearch(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error) {
	return ss.storage.ChSearch(ctx, searchQuery, params...)
}

func (ss searchService) PSearch(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error) {
	return ss.storage.PSearch(ctx, searchQuery, params...)
}
