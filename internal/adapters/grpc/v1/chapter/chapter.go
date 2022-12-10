package chapter_adapter

import (
	"context"
	"fmt"
	dto "read-only_web/internal/adapters/grpc/v1/chapter/dto"
	"read-only_web/internal/domain/entity"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

type chapterStorage struct {
	client pb.ChapterGRPCClient
}

func NewChapterStorage(client pb.ChapterGRPCClient) *chapterStorage {
	return &chapterStorage{client: client}
}

func (rs *chapterStorage) GetOneChapter(ctx context.Context, chapterID uint64) (entity.Chapter, error) {
	req := &pb.GetOneChapterRequest{ID: chapterID}
	resp, err := rs.client.GetOne(ctx, req)
	if err != nil {
		return entity.Chapter{}, err
	}
	fmt.Println("GetOneChapterResp", resp.UpdatedAt)

	return dto.ChapterFromGetOneChapterResponse(resp), err
}

func (rs *chapterStorage) GetAllChapters(ctx context.Context, regulationID uint64) ([]entity.Chapter, error) {
	req := &pb.GetAllChaptersByRegulationIdRequest{ID: regulationID}
	resp, err := rs.client.GetAll(ctx, req)
	if err != nil {
		return nil, err
	}

	return dto.ChaptersFromGetAllChaptersByRegulationIdResponse(resp), err
}
