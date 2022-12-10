package paragraph_adapter

import (
	"context"
	dto "read-only_web/internal/adapters/grpc/v1/paragraph/dto"
	"read-only_web/internal/domain/entity"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

type paragraphStorage struct {
	client pb.ParagraphGRPCClient
}

func NewChapterStorage(client pb.ParagraphGRPCClient) *paragraphStorage {
	return &paragraphStorage{client: client}
}

func (rs *paragraphStorage) GetAll(ctx context.Context, chapterID uint64) ([]entity.Paragraph, error) {
	req := &pb.GetAllParagraphsByChapterIdRequest{ID: chapterID}
	resp, err := rs.client.GetAll(ctx, req)
	if err != nil {
		return nil, err
	}

	return dto.ParagraphsFromGetAllParagraphsByChapterIdResponse(resp), err
}
