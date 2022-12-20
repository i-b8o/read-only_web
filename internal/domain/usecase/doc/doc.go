package usecase_doc

import (
	"context"
	"read-only_web/internal/domain/entity"
	"strconv"

	"github.com/i-b8o/logging"
	"github.com/i-b8o/nonsense"
)

type DocService interface {
	GetOne(ctx context.Context, docID uint64) (entity.Doc, error)
}
type ChapterService interface {
	GetAllChapters(ctx context.Context, docID uint64) ([]entity.ChapterInfo, error)
}

type docUsecase struct {
	docService     DocService
	chapterService ChapterService
	logger         logging.Logger
}

func NewDocUsecase(docService DocService, chapterService ChapterService, logger logging.Logger) *docUsecase {
	return &docUsecase{docService: docService, chapterService: chapterService, logger: logger}
}

func (u docUsecase) GetDocumentRoot(ctx context.Context, stringID string) *entity.Doc {
	uint64ID, err := strconv.ParseUint(stringID, 10, 64)
	if err != nil {
		u.logger.Infof("error '%v' has occurred while GetDocumentRoot processing reguklationID: %s", err, stringID)
		return nil
	}
	doc, err := u.docService.GetOne(ctx, uint64ID)
	if err != nil {
		u.logger.Infof("error '%v' has occurred while GetDocumentRoot processing reguklationID: %s", err, stringID)
		return nil
	}

	doc.Name = nonsense.Capitalize(doc.Name)
	chapters, err := u.chapterService.GetAllChapters(ctx, uint64ID)
	if err != nil {
		u.logger.Infof("error '%v' has occurred while GetDocumentRoot processing reguklationID: %s", err, stringID)
		return nil
	}
	doc.Chapters = chapters
	return &doc
}
