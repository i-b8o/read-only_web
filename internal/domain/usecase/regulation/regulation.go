package usecase_regulation

import (
	"context"
	"read-only_web/internal/domain/entity"
	"strconv"

	"github.com/i-b8o/logging"
	"github.com/i-b8o/nonsense"
)

type RegulationService interface {
	GetOne(ctx context.Context, regulationID uint64) (entity.Regulation, error)
}
type ChapterService interface {
	GetAllChapters(ctx context.Context, regulationID uint64) ([]entity.ChapterInfo, error)
}

type regulationUsecase struct {
	regulationService RegulationService
	chapterService    ChapterService
	logger            logging.Logger
}

func NewRegulationUsecase(regulationService RegulationService, chapterService ChapterService, logger logging.Logger) *regulationUsecase {
	return &regulationUsecase{regulationService: regulationService, chapterService: chapterService, logger: logger}
}

func (u regulationUsecase) GetDocumentRoot(ctx context.Context, stringID string) entity.Regulation {
	uint64ID, err := strconv.ParseUint(stringID, 10, 64)
	if err != nil {
		u.logger.Error(err)
		return entity.Regulation{}
	}
	regulation, err := u.regulationService.GetOne(ctx, uint64ID)
	if err != nil {
		u.logger.Error(err)
		return entity.Regulation{}
	}

	regulation.Name = nonsense.Capitalize(regulation.Name)
	chapters, err := u.chapterService.GetAllChapters(ctx, uint64ID)
	if err != nil {
		u.logger.Error(err)
		return entity.Regulation{}
	}
	regulation.Chapters = chapters
	return regulation
}
