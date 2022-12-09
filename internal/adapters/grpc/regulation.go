package grpc_adapter

import (
	"context"
	adapters_grpc_dto "read-only_web/internal/adapters/grpc/dto"
	"read-only_web/internal/domain/entity"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

type readerStorage struct {
	client pb.ReaderGRPCClient
}

func NewReaderStorage(client pb.ReaderGRPCClient) *readerStorage {
	return &readerStorage{client: client}
}

func (rs *readerStorage) GetOne(ctx context.Context, regulationID uint64) (entity.Regulation, error) {
	req := &pb.GetOneRegulationRequest{ID: regulationID}
	resp, err := rs.client.GetOneRegulation(ctx, req)
	if err != nil {
		return entity.Regulation{}, err
	}

	return adapters_grpc_dto.RegulationFromGetOneRegulationResponse(resp), err
}

func (rs *readerStorage) GetOneChapter(ctx context.Context, chapterID uint64) (entity.Chapter, error) {
	req := &pb.GetOneChapterRequest{ID: chapterID}
	resp, err := rs.client.GetOneChapter(ctx, req)
	if err != nil {
		return entity.Chapter{}, err
	}

	return adapters_grpc_dto.RegulationFromGetOneRegulationResponse(resp), err
}
