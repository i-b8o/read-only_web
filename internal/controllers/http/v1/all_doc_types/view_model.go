package alldoctypes_controller

import (
	"context"
	"read-only_web/internal/domain/entity"
)

type AllDocTypeUsecase interface {
	GetAllDocTypes(ctx context.Context) []entity.DocType
}

type viewModelState struct {
	Title       string
	Description string
	Keywords    string
	DocTypes    []entity.DocType
}

type viewModel struct {
	allDocTypesUsecase AllDocTypeUsecase
}

func NewViewModel(allDocTypesUsecase AllDocTypeUsecase) *viewModel {
	return &viewModel{allDocTypesUsecase: allDocTypesUsecase}
}

func (vm viewModel) GetState(ctx context.Context) *viewModelState {
	docTypes := vm.allDocTypesUsecase.GetAllDocTypes(ctx)
	if docTypes == nil {
		return nil
	}
	// TODO hard coding
	s := viewModelState{
		Title:       "Все документы",
		Description: "перечень правил и инструкций по охране труда",
		Keywords:    "перечень правил, инструкции, охрана труда",
		DocTypes:    docTypes,
	}
	return &s
}
