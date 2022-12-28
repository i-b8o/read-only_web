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
	const sql = `SELECT name,title,description,keywords,rev FROM "doc" WHERE id = $1 LIMIT 1`
	row := rs.client.QueryRow(ctx, sql, docID)

	doc := entity.Doc{}
	err := row.Scan(&doc.Name, &doc.Title, &doc.Description, &doc.Keywords, &doc.Rev)
	if err != nil {
		return entity.Doc{}, err
	}

	return doc, nil
}

func (s *docStorage) GetBySubtype(ctx context.Context, subTypeID uint64) ([]entity.Doc, error) {
	const sql = `select id, name from doc where subtype_id = $1;`

	var docs []entity.Doc

	rows, err := s.client.Query(ctx, sql, subTypeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		doc := entity.Doc{}
		if err = rows.Scan(
			&doc.ID, &doc.Name,
		); err != nil {
			return nil, err
		}

		docs = append(docs, doc)
	}
	return docs, nil
}
