package adapters_grpc_regulation_dto

import (
	"read-only_web/internal/domain/entity"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

func RegulationFromGetOneRegulationResponse(resp *pb.GetOneRegulationResponse) entity.Regulation {
	return entity.Regulation{Name: resp.Name, Abbreviation: resp.Abbreviation, Title: resp.Title}
}
