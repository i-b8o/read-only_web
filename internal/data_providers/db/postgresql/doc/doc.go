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

func (rs *docStorage) GetTypes(ctx context.Context) ([]string, error) {
	const sql = `select distinct(type) from doc order by type`

	var docs []entity.Doc

	rows, err := cs.client.Query(ctx, sql, docID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		chapter := entity.ChapterInfo{}
		if err = rows.Scan(
			&chapter.ID, &chapter.Name, &chapter.Num, &chapter.OrderNum,
		); err != nil {
			return nil, err
		}

		docs = append(docs, chapter)
	}
	return docs, nil
}

func (rs *docStorage) Get(ctx context.Context, docID uint64) (entity.Doc, error) {
	const sql = `SELECT name,title,description,keywords FROM "doc" WHERE id = $1 LIMIT 1`
	row := rs.client.QueryRow(ctx, sql, docID)

	doc := entity.Doc{}
	err := row.Scan(&doc.Name, &doc.Title, &doc.Description, &doc.Keywords)
	if err != nil {
		return entity.Doc{}, err
	}

	return doc, nil
}
