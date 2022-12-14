package service

import (
	"context"
	"read-only_web/internal/domain/entity"
)

type RegulationStorage interface {
	Get(ctx context.Context, regulationID uint64) (entity.Regulation, error)
}
type regulationService struct {
	storage RegulationStorage
}

func NewRegulationService(storage RegulationStorage) *regulationService {
	return &regulationService{storage: storage}
}

func (s *regulationService) GetOne(ctx context.Context, regulationID uint64) (entity.Regulation, error) {
	return s.storage.Get(ctx, regulationID)
}
