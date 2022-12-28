package usecase_doc

import (
	"context"
	"read-only_web/internal/domain/entity"

	"github.com/i-b8o/logging"
)

type DocService interface {
	GetOne(ctx context.Context, docID uint64) entity.Doc
	GetBySubtype(ctx context.Context, subtypeID uint64) []entity.Doc
}
type ChapterService interface {
	GetAll(ctx context.Context, docID uint64) []entity.ChapterInfo
}

type docUsecase struct {
	docService     DocService
	chapterService ChapterService
	logger         logging.Logger
}

func NewDocUsecase(docService DocService, chapterService ChapterService, logger logging.Logger) *docUsecase {
	return &docUsecase{docService: docService, chapterService: chapterService, logger: logger}
}

func (u docUsecase) GetDocumentRoot(ctx context.Context, docID uint64) *entity.Doc {
	// get a doc
	doc := u.docService.GetOne(ctx, docID)

	// get chapters for a doc
	chapters := u.chapterService.GetAll(ctx, docID)

	// insert the chapters into the doc
	doc.Chapters = chapters
	return &doc
}

func (u docUsecase) GetBySubtype(ctx context.Context, subtypeID uint64) []entity.Doc {
	return u.docService.GetBySubtype(ctx, subtypeID)
}
