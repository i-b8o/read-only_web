package service

import (
	"context"
	"read-only_web/internal/domain/entity"

	"github.com/i-b8o/logging"
)

type SubTypeStorage interface {
	GetAll(ctx context.Context, typeID uint64) ([]entity.DocSubType, error)
}
type subTypeService struct {
	storage SubTypeStorage
	logger  logging.Logger
}

func NewSubTypeService(storage SubTypeStorage, logger logging.Logger) *subTypeService {
	return &subTypeService{storage: storage, logger: logger}
}

func (s *subTypeService) GetAll(ctx context.Context, typeID uint64) []entity.DocSubType {
	subTypes, err := s.storage.GetAll(ctx, typeID)
	if err != nil {
		s.logger.Infof("error '%v' has occurred while GetAll processing", err)
		return []entity.DocSubType{}
	}
	return subTypes
}
