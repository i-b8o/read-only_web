package subtypes_controller

import (
	"context"
	"read-only_web/internal/domain/entity"
	"strconv"
)

type SubTypeUsecase interface {
	GetSubTypes(ctx context.Context, typeID uint64) ([]entity.DocSubType, []entity.DocType)
}

type viewModelState struct {
	Title       string
	Description string
	Keywords    string
	TypeID      uint64
	DocTypes    []entity.DocType
	DocSubTypes []entity.DocSubType
}

type viewModel struct {
	allDocTypesUsecase SubTypeUsecase
}

func NewViewModel(allDocTypesUsecase SubTypeUsecase) *viewModel {
	return &viewModel{allDocTypesUsecase: allDocTypesUsecase}
}

func (vm viewModel) GetState(ctx context.Context, id string) *viewModelState {
	// type-conversion then validate id is a positive num
	uint64ID, err := strconv.ParseUint(id, 10, 64)
	if err != nil || uint64ID <= 0 {
		return nil
	}
	docSubTypes, docTypes := vm.allDocTypesUsecase.GetSubTypes(ctx, uint64ID)
	if docTypes == nil {
		return nil
	}
	// TODO hard coding
	s := viewModelState{
		Title:       "Все документы",
		Description: "перечень правил и инструкций по охране труда",
		Keywords:    "перечень правил, инструкции, охрана труда",
		TypeID:      uint64ID,
		DocTypes:    docTypes,
		DocSubTypes: docSubTypes,
	}
	return &s
}
