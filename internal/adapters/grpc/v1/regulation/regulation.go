package grpc_adapter

import (
	"context"
	dto "read-only_web/internal/adapters/grpc/v1/regulation/dto"
	"read-only_web/internal/domain/entity"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

type regulationStorage struct {
	client pb.RegulationGRPCClient
}

func NewRegulationStorage(client pb.RegulationGRPCClient) *regulationStorage {
	return &regulationStorage{client: client}
}

func (rs *regulationStorage) GetOne(ctx context.Context, regulationID uint64) (entity.Regulation, error) {
	req := &pb.GetOneRegulationRequest{ID: regulationID}
	resp, err := rs.client.GetOne(ctx, req)
	if err != nil {
		return entity.Regulation{}, err
	}

	return dto.RegulationFromGetOneRegulationResponse(resp), err
}
