package service

import (
	"context"
	"read-only_web/internal/domain/entity"
)

type ChapterStorage interface {
	GetOneChapter(ctx context.Context, chapterID uint64) (entity.Chapter, error)
	GetAllChapters(ctx context.Context, regulationID uint64) ([]entity.Chapter, error)
}
type chapterService struct {
	storage ChapterStorage
}

func NewChapterService(storage ChapterStorage) *chapterService {
	return &chapterService{storage: storage}
}

func (s *chapterService) GetOneChapter(ctx context.Context, chapterID uint64) (entity.Chapter, error) {
	return s.storage.GetOneChapter(ctx, chapterID)
}
func (s *chapterService) GetAllChapters(ctx context.Context, regulationID uint64) ([]entity.Chapter, error) {
	return s.storage.GetAllChapters(ctx, regulationID)
}
