package service

import (
	"context"
	"read-only_web/internal/domain/entity"
)

type ChapterStorage interface {
	Get(ctx context.Context, chapterID uint64) (entity.Chapter, error)
	GetAll(ctx context.Context, regulationID uint64) ([]entity.ChapterInfo, error)
}
type chapterService struct {
	storage ChapterStorage
}

func NewChapterService(storage ChapterStorage) *chapterService {
	return &chapterService{storage: storage}
}

func (s *chapterService) GetOneChapter(ctx context.Context, chapterID uint64) (entity.Chapter, error) {
	return s.storage.Get(ctx, chapterID)
}
func (s *chapterService) GetAllChapters(ctx context.Context, regulationID uint64) ([]entity.ChapterInfo, error) {
	return s.storage.GetAll(ctx, regulationID)
}
