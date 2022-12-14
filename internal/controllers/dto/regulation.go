package dto

import (
	"fmt"
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
