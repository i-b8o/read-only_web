package chapter_controller

import (
	"context"
	"html/template"
	"read-only_web/internal/domain/entity"
	"strconv"
)

type ChapterUsecase interface {
	GetChapter(ctx context.Context, chapterID uint64) (*entity.Doc, *entity.Chapter)
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
	ChapterID   uint64
	Name        string
	Title       string
	Description string
	Keywords    string
	Rev         string
	Prev        entity.ChapterInfo
	Next        entity.ChapterInfo
	Num         string
	Paragraphs  []paragraph
	Chapters    []entity.ChapterInfo
}

type viewModel struct {
	chapterUsecase ChapterUsecase
}

func NewViewModel(chapterUsecase ChapterUsecase) *viewModel {
	return &viewModel{chapterUsecase: chapterUsecase}
}

func (vm viewModel) GetState(ctx context.Context, id string) *viewModelState {
	// type-conversion then validate id is a positive num
	uint64ID, err := strconv.ParseUint(id, 10, 64)
	if err != nil || uint64ID <= 0 {
		return nil
	}

	doc, chapter := vm.chapterUsecase.GetChapter(ctx, uint64ID)
	if doc == nil || chapter == nil {
		return nil
	}
	var prevChapter, nextChapter entity.ChapterInfo

	// the chapter order num starts from 1 (not 0)
	if chapter.OrderNum > 1 {
		prevChapter = doc.Chapters[chapter.OrderNum-2]
	}
	if int(chapter.OrderNum) < len(doc.Chapters) {
		nextChapter = doc.Chapters[chapter.OrderNum]
	}

	var paragraphs []paragraph
	for _, p := range chapter.Paragraphs {
		paragraph := paragraph{ID: p.ID, Num: p.Num, Class: p.Class, Content: template.HTML(p.Content), ChapterID: p.ChapterID, IsTable: p.IsTable}
		paragraphs = append(paragraphs, paragraph)
	}

	s := viewModelState{
		ChapterID:   chapter.ID,
		Title:       chapter.Title,
		Description: chapter.Description,
		Keywords:    chapter.Keywords,
		Name:        chapter.Name,
		Rev:         chapter.Rev,
		Num:         chapter.Num,
		Chapters:    doc.Chapters,
		Prev:        prevChapter,
		Next:        nextChapter,
		Paragraphs:  paragraphs,
	}
	return &s
}
