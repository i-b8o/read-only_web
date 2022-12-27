package service

import (
	"context"
	"read-only_web/internal/domain/entity"

	"github.com/i-b8o/logging"
)

type ChapterStorage interface {
	Get(ctx context.Context, chapterID uint64) (entity.Chapter, error)
	GetAll(ctx context.Context, docID uint64) ([]entity.ChapterInfo, error)
}
type chapterService struct {
	storage ChapterStorage
	logger  logging.Logger
}

func NewChapterService(storage ChapterStorage, logger logging.Logger) *chapterService {
	return &chapterService{storage: storage, logger: logger}
}

func (s *chapterService) GetOne(ctx context.Context, chapterID uint64) entity.Chapter {
	chapter, err := s.storage.Get(ctx, chapterID)
	if err != nil {
		s.logger.Infof("error '%v' has occurred while GetOne processing chapterID: %s", err, chapterID)
		return entity.Chapter{}
	}
	return chapter
}

func (s *chapterService) GetAll(ctx context.Context, docID uint64) []entity.ChapterInfo {
	chapters, err := s.storage.GetAll(ctx, docID)
	if err != nil {
		s.logger.Infof("error '%v' has occurred while GetAll processing docID: %s", err, docID)
		return []entity.ChapterInfo{}
	}
	return chapters
}
