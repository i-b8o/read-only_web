package service

import (
	"context"
	"read-only_web/internal/domain/entity"
)

type DocStorage interface {
	Get(ctx context.Context, docID uint64) (entity.Doc, error)
}
type docService struct {
	storage DocStorage
}

func NewDocService(storage DocStorage) *docService {
	return &docService{storage: storage}
}

func (s *docService) GetOne(ctx context.Context, docID uint64) (entity.Doc, error) {
	return s.storage.Get(ctx, docID)
}
