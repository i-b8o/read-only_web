package usecase_chapter

import (
	"context"
	"read-only_web/internal/domain/entity"
	"strconv"

	"github.com/i-b8o/logging"
)

type ChapterService interface {
	GetOneChapter(ctx context.Context, chapterID uint64) (entity.Chapter, error)
	GetAllChapters(ctx context.Context, docID uint64) ([]entity.ChapterInfo, error)
}

type ParagraphService interface {
	GetAll(ctx context.Context, chapterID uint64) ([]entity.Paragraph, error)
}

type DocService interface {
	GetOne(ctx context.Context, docID uint64) (entity.Doc, error)
}

type chapterUsecase struct {
	chapterService   ChapterService
	paragraphService ParagraphService
	docService       DocService
	logger           logging.Logger
}

func NewChapterUsecase(chapterService ChapterService, paragraphService ParagraphService, docService DocService, logger logging.Logger) *chapterUsecase {
	return &chapterUsecase{chapterService: chapterService, paragraphService: paragraphService, docService: docService, logger: logger}
}

func (u chapterUsecase) GetChapter(ctx context.Context, chapterID string) (*entity.Doc, *entity.Chapter) {
	uint64ID, err := strconv.ParseUint(chapterID, 10, 64)
	if err != nil {
		u.logger.Infof("error '%v' has occurred while GetChapter processing chapterID: %s", err, chapterID)
		return nil, nil
	}

	chapter, err := u.chapterService.GetOneChapter(ctx, uint64ID)
	if err != nil {
		u.logger.Infof("error '%v' has occurred while GetChapter processing chapterID: %s", err, chapterID)
		return nil, nil
	}

	chapter.Paragraphs, err = u.paragraphService.GetAll(ctx, uint64ID)
	if err != nil {
		u.logger.Infof("error '%v' has occurred while GetChapter processing chapterID: %s", err, chapterID)
		return nil, nil
	}
	doc, err := u.docService.GetOne(ctx, chapter.DocID)
	if err != nil {
		u.logger.Infof("error '%v' has occurred while GetChapter processing chapterID: %s", err, chapterID)
		return nil, nil
	}
	chapters, err := u.chapterService.GetAllChapters(ctx, chapter.DocID)
	if err != nil {
		u.logger.Infof("error '%v' has occurred while GetChapter processing chapterID: %s", err, chapterID)
		return nil, nil
	}
	doc.Chapters = chapters
	return &doc, &chapter
}
