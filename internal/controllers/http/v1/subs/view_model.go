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
	Header           string
	Title            string
	Description      string
	Keywords         string
	CurrentDocTypeID uint64
	DocTypes         *[]entity.DocType
	DocSubTypes      []entity.DocSubType
}

type viewModel struct {
	allDocTypesUsecase SubTypeUsecase
	docUsecase         DocUsecase
}

func NewViewModel(allDocTypesUsecase SubTypeUsecase, docUsecase DocUsecase) *viewModel {
	return &viewModel{allDocTypesUsecase: allDocTypesUsecase, docUsecase: docUsecase}
}

func (vm viewModel) GetState(ctx context.Context, typeID string) *viewModelState {
	// type-conversion then validate id is a positive num
	uint64ID, err := strconv.ParseUint(typeID, 10, 64)
	if err != nil || uint64ID <= 0 {
		return nil
	}
	docSubTypes, docTypes := vm.allDocTypesUsecase.GetSubTypes(ctx, uint64ID)
	if docTypes == nil {
		return nil
	}
	idx := slices.IndexFunc(docTypes, func(dst entity.DocType) bool { return dst.ID == uint64ID })

	// TODO hard coding
	s := viewModelState{
		Title:            "",
		Description:      "",
		Keywords:         "",
		Header:           docTypes[idx].Name,
		CurrentDocTypeID: uint64ID,
		DocTypes:         &docTypes,
		DocSubTypes:      docSubTypes,
	}
	return &s
}
