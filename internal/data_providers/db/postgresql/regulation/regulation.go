package regulation_provider

import (
	"context"

	"read-only_web/internal/domain/entity"
	client "read-only_web/pkg/client/postgresql"
)

type regulationStorage struct {
	client client.PostgreSQLClient
}

func NewRegulationStorage(client client.PostgreSQLClient) *regulationStorage {
	return &regulationStorage{client: client}
}

func (rs *regulationStorage) Get(ctx context.Context, regulationID uint64) (entity.Regulation, error) {
	const sql = `SELECT name,abbreviation,header,title,meta,keywords FROM "regulation" WHERE id = $1 LIMIT 1`
	row := rs.client.QueryRow(ctx, sql, regulationID)

	regulation := entity.Regulation{}
	err := row.Scan(&regulation.Name, &regulation.Abbreviation, &regulation.Header, &regulation.Title, &regulation.Meta, &regulation.Keywords)
	if err != nil {
		return entity.Regulation{}, err
	}

	return regulation, nil
}
