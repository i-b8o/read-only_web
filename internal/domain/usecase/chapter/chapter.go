package usecase_chapter

import (
	"context"
	"read-only_web/internal/domain/entity"
)

type ChapterService interface {
	GetOne(ctx context.Context, chapterID uint64) entity.Chapter
	GetAll(ctx context.Context, docID uint64) []entity.ChapterInfo
}

type ParagraphService interface {
	GetAll(ctx context.Context, chapterID uint64) []entity.Paragraph
}

type DocService interface {
	GetOne(ctx context.Context, docID uint64) entity.Doc
}

type chapterUsecase struct {
	chapterService   ChapterService
	paragraphService ParagraphService
	docService       DocService
}

func NewChapterUsecase(chapterService ChapterService, paragraphService ParagraphService, docService DocService) *chapterUsecase {
	return &chapterUsecase{chapterService: chapterService, paragraphService: paragraphService, docService: docService}
}

func (u chapterUsecase) GetChapter(ctx context.Context, chapterID uint64) (*entity.Doc, *entity.Chapter) {
	// get a chapter
	chapter := u.chapterService.GetOne(ctx, chapterID)

	// get a paragraphs for the chapter
	chapter.Paragraphs = u.paragraphService.GetAll(ctx, chapterID)

	// get a doc for the chapter
	doc := u.docService.GetOne(ctx, chapter.DocID)

	// get a chapters for the doc
	chapters := u.chapterService.GetAll(ctx, chapter.DocID)

	// insert the chapters into the doc
	doc.Chapters = chapters

	return &doc, &chapter
}
