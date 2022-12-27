package postgresql

import (
	"context"
	"read-only_web/internal/domain/entity"
	client "read-only_web/pkg/client/postgresql"
)

type typeStorage struct {
	client client.PostgreSQLClient
}

func NewTypeStorage(client client.PostgreSQLClient) *typeStorage {
	return &typeStorage{client: client}
}

func (ds *typeStorage) GetAll(ctx context.Context) ([]entity.DocType, error) {
	const sql = `select id, name from type order by type`

	var docTypes []entity.DocType

	rows, err := ds.client.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		docType := entity.DocType{}
		if err = rows.Scan(
			&docType.ID, &docType.Name,
		); err != nil {
			return nil, err
		}

		docTypes = append(docTypes, docType)
	}
	return docTypes, nil
}
