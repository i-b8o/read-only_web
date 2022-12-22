package doc_provider

import (
	"context"

	"read-only_web/internal/domain/entity"
	client "read-only_web/pkg/client/postgresql"
)

type docStorage struct {
	client client.PostgreSQLClient
}

func NewDocStorage(client client.PostgreSQLClient) *docStorage {
	return &docStorage{client: client}
}

func (rs *docStorage) Get(ctx context.Context, docID uint64) (entity.Doc, error) {
	const sql = `SELECT name,abbreviation,header,title,description FROM "doc" WHERE id = $1 LIMIT 1`
	row := rs.client.QueryRow(ctx, sql, docID)

	doc := entity.Doc{}
	err := row.Scan(&doc.Name, &doc.Abbreviation, &doc.Header, &doc.Title, &doc.Description)
	if err != nil {
		return entity.Doc{}, err
	}

	return doc, nil
}
