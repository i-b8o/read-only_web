package postgresql

import (
	"context"
	client "read-only_web/pkg/client/postgresql"
)

type subTypeDocStorage struct {
	client client.PostgreSQLClient
}

func NewSubTypeDocStorage(client client.PostgreSQLClient) *subTypeDocStorage {
	return &subTypeDocStorage{client: client}
}

func (s *subTypeDocStorage) GetAll(ctx context.Context, subtypeID uint64) ([]uint64, error) {
	const sql = `SELECT doc_id FROM subtype_doc WHERE subtype_id=$1`

	var docsIDs []uint64

	rows, err := s.client.Query(ctx, sql, subtypeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var docID uint64
		if err = rows.Scan(
			&docID,
		); err != nil {
			return nil, err
		}

		docsIDs = append(docsIDs, docID)
	}
	return docsIDs, nil
}
