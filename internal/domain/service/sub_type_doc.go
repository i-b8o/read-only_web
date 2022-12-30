package service

import (
	"context"

	"github.com/i-b8o/logging"
)

type SubTypeDocStorage interface {
	GetAll(ctx context.Context, subtypeID uint64) ([]uint64, error)
}
type subTypeDocService struct {
	storage SubTypeDocStorage
	logger  logging.Logger
}

func NewSubTypeDocService(storage SubTypeDocStorage, logger logging.Logger) *subTypeDocService {
	return &subTypeDocService{storage: storage, logger: logger}
}

func (s *subTypeDocService) GetAll(ctx context.Context, subtypeID uint64) []uint64 {
	docsIDs, err := s.storage.GetAll(ctx, subtypeID)
	if err != nil {
		s.logger.Infof("error '%v' has occurred while GetAll processing", err)
		return nil
	}
	return docsIDs
}
