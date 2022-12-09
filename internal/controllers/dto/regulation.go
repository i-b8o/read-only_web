package dto

import (
	"fmt"
	"read-only_web/internal/domain/entity"
)

type CreateRegulationRequestDTO struct {
	RegulationName string `json:"regulation_name"`
	Abbreviation   string `json:"abbreviation"`
}

func (dto *CreateRegulationRequestDTO) Validate() error {
	if dto.RegulationName == "" {
		return fmt.Errorf("missing regulation name")
	}

	return nil
}

type GetFullRegulationRequestDTO struct {
	RegulationID uint64 `json:"regulation_id"`
}

type GetFullRegulationJSONResponseDTO struct {
	Regulation entity.Regulation `json:"regulation"`
}

type GetFullRegulationDartResponseDTO struct {
	Regulation string `json:"regulation"`
}
