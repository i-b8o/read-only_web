package doc_controller

import (
	"context"
	"read-only_web/internal/domain/entity"
	"strconv"
)

type DocUsecase interface {
	GetDocumentRoot(ctx context.Context, stringID string) *entity.Doc
}

type viewModelState struct {
	Abbreviation string
	Header       *string
	Title        *string
	Description  *string
	Keywords     *string
	Name         string
	Chapters     []entity.ChapterInfo
}

type viewModel struct {
	docUsecase DocUsecase
}

func NewViewModel(docUsecase DocUsecase) *viewModel {
	return &viewModel{docUsecase: docUsecase}
}

func (vm viewModel) GetState(ctx context.Context, id string) *viewModelState {
	// validate id is a positive num
	n, err := strconv.ParseFloat(id, 64)
	if err != nil || n <= 0 {
		return nil
	}

	doc := vm.docUsecase.GetDocumentRoot(ctx, id)
	if doc == nil {
		return nil
	}
	s := viewModelState{
		Abbreviation: doc.Abbreviation,
		Header:       doc.Header,
		Title:        &doc.Name,
		Description:  doc.Description,
		Keywords:     doc.Keywords,
		Name:         doc.Name,
		Chapters:     doc.Chapters,
	}
	return &s
}
