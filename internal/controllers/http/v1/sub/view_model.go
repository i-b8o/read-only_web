package subtypes_controller

import (
	"context"
	"read-only_web/internal/domain/entity"
	"strconv"

	"golang.org/x/exp/slices"
)

type SubTypeUsecase interface {
	GetSubTypes(ctx context.Context, typeID uint64) ([]entity.DocSubType, []entity.DocType)
}

type DocUsecase interface {
	GetBySubtype(ctx context.Context, subtypeID uint64) []entity.Doc
}

type viewModelState struct {
	Header               string
	Title                string
	Description          string
	Keywords             string
	CurrentDocSubTypesID uint64
	CurrentDocTypesID    uint64
	DocSubTypes          []entity.DocSubType
	Docs                 *[]entity.Doc
}

type viewModel struct {
	allDocTypesUsecase SubTypeUsecase
	docUsecase         DocUsecase
}

func NewViewModel(allDocTypesUsecase SubTypeUsecase, docUsecase DocUsecase) *viewModel {
	return &viewModel{allDocTypesUsecase: allDocTypesUsecase, docUsecase: docUsecase}
}

func (vm viewModel) GetState(ctx context.Context, typeID, subtypeID string) *viewModelState {
	// type-conversion then validate id is a positive num
	uint64subtypeID, err := strconv.ParseUint(subtypeID, 10, 64)
	if err != nil || uint64subtypeID <= 0 {
		return nil
	}
	uint64typeID, err := strconv.ParseUint(typeID, 10, 64)
	if err != nil || uint64typeID <= 0 {
		return nil
	}
	docSubTypes, docTypes := vm.allDocTypesUsecase.GetSubTypes(ctx, uint64typeID)
	if docTypes == nil {
		return nil
	}
	docs := vm.docUsecase.GetBySubtype(ctx, uint64subtypeID)
	if docs == nil {
		return nil
	}
	idx := slices.IndexFunc(docSubTypes, func(dst entity.DocSubType) bool { return dst.ID == uint64subtypeID })
	// TODO empty title description keywords
	s := viewModelState{
		Title:                "",
		Description:          "",
		Keywords:             "",
		Header:               docSubTypes[idx].Name,
		CurrentDocSubTypesID: uint64subtypeID,
		CurrentDocTypesID:    uint64typeID,
		DocSubTypes:          docSubTypes,
		Docs:                 &docs,
	}
	return &s
}
