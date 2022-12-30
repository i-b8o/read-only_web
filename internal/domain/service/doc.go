package service

import (
	"context"
	"read-only_web/internal/domain/entity"

	"github.com/i-b8o/logging"
	"github.com/i-b8o/nonsense"
)

type DocStorage interface {
	Get(ctx context.Context, docID uint64) (entity.Doc, error)
}
type docService struct {
	storage DocStorage
	logger  logging.Logger
}

func NewDocService(storage DocStorage, logger logging.Logger) *docService {
	return &docService{storage: storage, logger: logger}
}

func (s *docService) GetOne(ctx context.Context, docID uint64) entity.Doc {
	doc, err := s.storage.Get(ctx, docID)
	if err != nil {
		s.logger.Infof("error '%v' has occurred while GetOne processing docID: %s", err, docID)
		return entity.Doc{}
	}
	doc.Name = nonsense.Capitalize(doc.Name)
	return doc
}
