package regulation_controller

import (
	"context"
	"read-only_web/internal/domain/entity"
)

type RegulationUsecase interface {
	GetDocumentRoot(ctx context.Context, stringID string) entity.Regulation
}

type viewModelState struct {
	abbreviation string
	header       *string
	title        *string
	meta         *string
	keywords     *string
	name         string
	chapters     []entity.Chapter
}

type viewModel struct {
	regulationUsecase RegulationUsecase
}

func NewViewModel(regulationUsecase RegulationUsecase) *viewModel {
	return &viewModel{regulationUsecase: regulationUsecase}
}

func (vm viewModel) GetState(ctx context.Context, id string) *viewModelState {
	regulation := vm.regulationUsecase.GetDocumentRoot(ctx, id)
	s := viewModelState{abbreviation: regulation.Abbreviation, header: regulation.Header, title: regulation.Title, meta: regulation.Meta, keywords: regulation.Keywords, name: regulation.Name, chapters: regulation.Chapters}
	return &s
}
