package chapter_controller

import (
	"context"
	"html/template"
	"read-only_web/internal/domain/entity"
)

type ChapterUsecase interface {
	GetChapter(ctx context.Context, chapterID string) (entity.Regulation, entity.Chapter)
}

type paragraph struct {
	ID        uint64        `json:"id"`
	Num       uint64        `json:"num"`
	IsTable   bool          `json:"is_table"`
	Class     string        `json:"class,omitempty"`
	Content   template.HTML `json:"content,omitempty"`
	ChapterID uint64        `json:"chapterid,omitempty"`
}

type viewModelState struct {
	// regulations  []entity.Regulation
	chapterID    uint64
	name         string
	abbreviation string
	// header       *string
	// title        *string
	// meta         *string
	// keywords     *string
	prev       entity.ChapterInfo
	next       entity.ChapterInfo
	num        string
	paragraphs []paragraph
	chapters   []entity.ChapterInfo
	updated    string
}

type viewModel struct {
	chapterUsecase ChapterUsecase
}

func NewViewModel(chapterUsecase ChapterUsecase) *viewModel {
	return &viewModel{chapterUsecase: chapterUsecase}
}

func (vm viewModel) GetState(ctx context.Context, id string) *viewModelState {
	regulation, chapter := vm.chapterUsecase.GetChapter(ctx, id)
	if regulation.IsEmpty() || chapter.IsEmpty() {
		return nil
	}
	var prevChapter, nextChapter entity.ChapterInfo

	// the chapter order num starts from 1 (not 0)
	if chapter.OrderNum > 1 {
		prevChapter = regulation.Chapters[chapter.OrderNum-2]
	}
	if int(chapter.OrderNum) < len(regulation.Chapters) {
		nextChapter = regulation.Chapters[chapter.OrderNum]
	}

	var paragraphs []paragraph
	for _, p := range chapter.Paragraphs {
		templParagraph := paragraph{ID: p.ID, Num: p.Num, Class: p.Class, Content: template.HTML(p.Content), ChapterID: p.ChapterID, IsTable: p.IsTable}
		paragraphs = append(paragraphs, templParagraph)
	}

	s := viewModelState{
		chapterID:    chapter.ID,
		abbreviation: regulation.Abbreviation,
		name:         chapter.Name,
		num:          chapter.Num,
		chapters:     regulation.Chapters,
		prev:         prevChapter,
		next:         nextChapter,
		paragraphs:   paragraphs,
		updated:      chapter.UpdatedAt.Format("02.01.2006"),
	}
	return &s
}