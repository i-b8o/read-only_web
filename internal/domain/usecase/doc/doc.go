package usecase_doc

import (
	"context"
	"read-only_web/internal/domain/entity"
)

type DocService interface {
	GetOne(ctx context.Context, docID uint64) entity.Doc
}
type ChapterService interface {
	GetAll(ctx context.Context, docID uint64) []entity.ChapterInfo
}
type SubTypeDocService interface {
	GetAll(ctx context.Context, subtypeID uint64) []uint64
}
type docUsecase struct {
	docService        DocService
	chapterService    ChapterService
	subTypeDocService SubTypeDocService
}

func NewDocUsecase(docService DocService, chapterService ChapterService, subTypeDocService SubTypeDocService) *docUsecase {
	return &docUsecase{docService: docService, chapterService: chapterService, subTypeDocService: subTypeDocService}
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

func (u docUsecase) GetBySubtype(ctx context.Context, subtypeID uint64) (docs []entity.Doc) {
	docsIDs := u.subTypeDocService.GetAll(ctx, subtypeID)

	for _, id := range docsIDs {
		doc := u.docService.GetOne(ctx, id)
		docs = append(docs, doc)
	}
	return docs
}
