package service

import (
	"context"
	"read-only_web/internal/domain/entity"

	"github.com/i-b8o/logging"
)

type ParagraphStorage interface {
	GetAll(ctx context.Context, chapterID uint64) ([]entity.Paragraph, error)
}
type paragraphService struct {
	storage ParagraphStorage
	logger  logging.Logger
}

func NewParagraphService(storage ParagraphStorage, logger logging.Logger) *paragraphService {
	return &paragraphService{storage: storage, logger: logger}
}

func (s *paragraphService) GetAll(ctx context.Context, chapterID uint64) []entity.Paragraph {
	paragraphs, err := s.storage.GetAll(ctx, chapterID)
	if err != nil {
		s.logger.Infof("error '%v' has occurred while GetAll processing chapterID: %s", err, chapterID)
		return []entity.Paragraph{}
	}
	return paragraphs
}
