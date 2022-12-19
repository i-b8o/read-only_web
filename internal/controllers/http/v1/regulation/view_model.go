package regulation_controller

import (
	"context"
	"read-only_web/internal/domain/entity"
)

type RegulationUsecase interface {
	GetDocumentRoot(ctx context.Context, stringID string) entity.Regulation
}

type viewModelState struct {
	Abbreviation string
	Header       *string
	Title        *string
	Meta         *string
	Keywords     *string
	Name         string
	Chapters     []entity.ChapterInfo
}

type viewModel struct {
	regulationUsecase RegulationUsecase
}

func NewViewModel(regulationUsecase RegulationUsecase) *viewModel {
	return &viewModel{regulationUsecase: regulationUsecase}
}

func (vm viewModel) GetState(ctx context.Context, id string) *viewModelState {
	regulation := vm.regulationUsecase.GetDocumentRoot(ctx, id)
	if regulation.IsEmpty() {
		return nil
	}
	s := viewModelState{Abbreviation: regulation.Abbreviation, Header: regulation.Header, Title: &regulation.Name, Meta: regulation.Meta, Keywords: regulation.Keywords, Name: regulation.Name, Chapters: regulation.Chapters}
	return &s
}
