package usecase_chapter

import (
	"context"
	"fmt"
	"read-only_web/internal/domain/entity"
	"strconv"

	"github.com/i-b8o/logging"
)

type ChapterService interface {
	GetOneChapter(ctx context.Context, chapterID uint64) (entity.Chapter, error)
	GetAllChapters(ctx context.Context, regulationID uint64) ([]entity.ChapterInfo, error)
}

type ParagraphService interface {
	GetAll(ctx context.Context, chapterID uint64) ([]entity.Paragraph, error)
}

type RegulationService interface {
	GetOne(ctx context.Context, regulationID uint64) (entity.Regulation, error)
}

type chapterUsecase struct {
	chapterService    ChapterService
	paragraphService  ParagraphService
	regulationService RegulationService
	logger            logging.Logger
}

func NewChapterUsecase(chapterService ChapterService, paragraphService ParagraphService, regulationService RegulationService, logger logging.Logger) *chapterUsecase {
	return &chapterUsecase{chapterService: chapterService, paragraphService: paragraphService, regulationService: regulationService, logger: logger}
}

// TODO do not send an error when a chapter does not exist
func (u chapterUsecase) GetChapter(ctx context.Context, chapterID string) (entity.Regulation, entity.Chapter) {
	uint64ID, err := strconv.ParseUint(chapterID, 10, 64)
	if err != nil {
		u.logger.Error(err)
		return entity.Regulation{}, entity.Chapter{}
	}

	chapter, err := u.chapterService.GetOneChapter(ctx, uint64ID)
	if err != nil {
		fmt.Println("chapter")
		u.logger.Error(err)
		return entity.Regulation{}, entity.Chapter{}
	}

	chapter.Paragraphs, err = u.paragraphService.GetAll(ctx, uint64ID)
	if err != nil {
		fmt.Println("paragraphs")
		u.logger.Error(err)
		return entity.Regulation{}, entity.Chapter{}
	}
	regulation, err := u.regulationService.GetOne(ctx, chapter.RegulationID)
	if err != nil {
		fmt.Println("regulation")
		u.logger.Error(err)
		return entity.Regulation{}, entity.Chapter{}
	}
	chapters, err := u.chapterService.GetAllChapters(ctx, chapter.RegulationID)
	if err != nil {
		fmt.Println("chapters")
		u.logger.Error(err)
		return entity.Regulation{}, entity.Chapter{}
	}
	regulation.Chapters = chapters
	return regulation, chapter
}
