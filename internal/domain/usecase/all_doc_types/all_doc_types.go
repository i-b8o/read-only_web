package usecase_all_doc_types

import (
	"context"
	"read-only_web/internal/domain/entity"
)

type TypeService interface {
	GetAll(ctx context.Context) []entity.DocType
}

type SubTypeService interface {
	GetAll(ctx context.Context, typeID uint64) []entity.DocSubType
}

type allTypesUsecase struct {
	typeService    TypeService
	subTypeService SubTypeService
}

func NewAllTypesUsecase(typeService TypeService, subTypeService SubTypeService) *allTypesUsecase {
	return &allTypesUsecase{typeService: typeService, subTypeService: subTypeService}
}

func (u allTypesUsecase) GetAllDocTypes(ctx context.Context) []entity.DocType {
	return u.typeService.GetAll(ctx)
}

func (u allTypesUsecase) GetSubTypes(ctx context.Context, typeID uint64) ([]entity.DocSubType, []entity.DocType) {
	types := u.typeService.GetAll(ctx)
	subtypes := u.subTypeService.GetAll(ctx, typeID)
	return subtypes, types
}
