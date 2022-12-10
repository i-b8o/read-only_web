package service

import (
	"context"
	"read-only_web/internal/domain/entity"
)

type ParagraphStorage interface {
	GetAll(ctx context.Context, chapterID uint64) ([]entity.Paragraph, error)
}
type paragraphService struct {
	storage ParagraphStorage
}

func NewParagraphService(storage ParagraphStorage) *paragraphService {
	return &paragraphService{storage: storage}
}

func (s *paragraphService) GetAll(ctx context.Context, chapterID uint64) ([]entity.Paragraph, error) {
	return s.storage.GetAll(ctx, chapterID)
}
