package dto

import (
	"fmt"
)

type CreateDocRequestDTO struct {
	DocName      string `json:"doc_name"`
	Abbreviation string `json:"abbreviation"`
}

func (dto *CreateDocRequestDTO) Validate() error {
	if dto.DocName == "" {
		return fmt.Errorf("missing doc name")
	}

	return nil
}
