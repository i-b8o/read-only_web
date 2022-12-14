package paragraph_provider

import (
	"context"
	"errors"
	"fmt"

	"read-only_web/internal/domain/entity"
	client "read-only_web/pkg/client/postgresql"

	"github.com/jackc/pgconn"
)

type paragraphStorage struct {
	client client.PostgreSQLClient
}

func NewParagraphStorage(client client.PostgreSQLClient) *paragraphStorage {
	return &paragraphStorage{client: client}
}

// GetAllById returns all paragraphs associated with the given chapter ID
func (ps *paragraphStorage) GetAll(ctx context.Context, chapterID uint64) ([]entity.Paragraph, error) {
	const sql = `SELECT paragraph_id, order_num, is_nft, is_table, has_links, class, content, c_id FROM "paragraph" WHERE c_id = $1 AND content!='-' ORDER BY order_num`

	var paragraphs []entity.Paragraph

	rows, err := ps.client.Query(ctx, sql, chapterID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return nil, fmt.Errorf("message: %s, code: %s", pgErr.Message, pgErr.Code)
		}
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		paragraph := entity.Paragraph{}
		if err = rows.Scan(
			&paragraph.ID, &paragraph.Num, &paragraph.IsNFT, &paragraph.IsTable, &paragraph.HasLinks, &paragraph.Class, &paragraph.Content, &paragraph.ChapterID,
		); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				return nil, fmt.Errorf("message: %s, code: %s", pgErr.Message, pgErr.Code)
			}
			return nil, err
		}

		paragraphs = append(paragraphs, paragraph)
	}

	return paragraphs, nil
}
