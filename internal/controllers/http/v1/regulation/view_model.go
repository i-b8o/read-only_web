package regulation_controller

import (
	"context"
	"fmt"
	"read-only_web/internal/domain/entity"
)

type RegulationUsecase interface {
	GetDocumentRoot(ctx context.Context, stringID string) entity.Regulation
}

type ViewModelState struct {
	Abbreviation string
	Header       *string
	Title        *string
	Meta         *string
	Keywords     *string
	Name         string
	Chapters     []entity.ChapterInfo
}

type ViewModel struct {
	regulationUsecase RegulationUsecase
}

func NewViewModel(regulationUsecase RegulationUsecase) *ViewModel {
	return &ViewModel{regulationUsecase: regulationUsecase}
}

func (vm ViewModel) GetState(ctx context.Context, id string) *ViewModelState {
	fmt.Println("zzzzzzzzzzzzzzzzzzzzzzzzzzz")
	regulation := vm.regulationUsecase.GetDocumentRoot(ctx, id)
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaa" + regulation.Abbreviation)
	if regulation.IsEmpty() {
		return nil
	}
	s := ViewModelState{Abbreviation: regulation.Abbreviation, Header: regulation.Header, Title: regulation.Title, Meta: regulation.Meta, Keywords: regulation.Keywords, Name: regulation.Name, Chapters: regulation.Chapters}
	return &s
}
