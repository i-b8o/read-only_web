package doc_controller

import (
	"context"
	"read-only_web/internal/domain/entity"
	"strconv"
)

type DocUsecase interface {
	GetDocumentRoot(ctx context.Context, docID uint64) *entity.Doc
}

type viewModelState struct {
	Title       string
	Description string
	Keywords    string
	Name        string
	Chapters    []entity.ChapterInfo
}

type viewModel struct {
	docUsecase DocUsecase
}

func NewViewModel(docUsecase DocUsecase) *viewModel {
	return &viewModel{docUsecase: docUsecase}
}

func (vm viewModel) GetState(ctx context.Context, id string) *viewModelState {
	// type-conversion then validate id is a positive num
	uint64ID, err := strconv.ParseUint(id, 10, 64)
	if err != nil || uint64ID <= 0 {
		return nil
	}

	doc := vm.docUsecase.GetDocumentRoot(ctx, uint64ID)
	if doc == nil {
		return nil
	}
	s := viewModelState{
		Title:       doc.Name,
		Description: doc.Description,
		Keywords:    doc.Keywords,
		Name:        doc.Name,
		Chapters:    doc.Chapters,
	}
	return &s
}

func (vm viewModel) GetDefaultState() *viewModelState {
	s := viewModelState{
		Title:       "Главная",
		Description: "Законодательство - законы и кодексы Российской Федерации. Полные тексты документов в последней редакции.",
		Keywords:    "Законодательство, законы, кодексы, федеральные законы, документы, Россия, РФ, налогообложение, налоги, трудовое, семейное, налоговое, административное, право,",
	}
	return &s
}
