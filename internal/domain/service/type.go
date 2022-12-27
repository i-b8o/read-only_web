package service

import (
	"context"
	"read-only_web/internal/domain/entity"

	"github.com/i-b8o/logging"
)

type TypeStorage interface {
	GetAll(ctx context.Context) ([]entity.DocType, error)
}
type typeService struct {
	storage TypeStorage
	logger  logging.Logger
}

func NewTypeService(storage TypeStorage, logger logging.Logger) *typeService {
	return &typeService{storage: storage, logger: logger}
}

func (s *typeService) GetAll(ctx context.Context) []entity.DocType {
	types, err := s.storage.GetAll(ctx)
	if err != nil {
		s.logger.Infof("error '%v' has occurred while GetAll processing", err)
		return []entity.DocType{}
	}
	return types
}
