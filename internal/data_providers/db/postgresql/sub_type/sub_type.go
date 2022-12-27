package postgresql

import (
	"context"
	"read-only_web/internal/domain/entity"
	client "read-only_web/pkg/client/postgresql"
)

type subTypeStorage struct {
	client client.PostgreSQLClient
}

func NewSubTypeStorage(client client.PostgreSQLClient) *subTypeStorage {
	return &subTypeStorage{client: client}
}

func (s *subTypeStorage) GetAll(ctx context.Context, typeID uint64) ([]entity.DocSubType, error) {
	const sql = `SELECT id, name FROM subtype WHERE type_id=$1`

	var docSubTypes []entity.DocSubType

	rows, err := s.client.Query(ctx, sql, typeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		docSubType := entity.DocSubType{}
		if err = rows.Scan(
			&docSubType.ID, &docSubType.Name,
		); err != nil {
			return nil, err
		}

		docSubTypes = append(docSubTypes, docSubType)
	}
	return docSubTypes, nil
}
