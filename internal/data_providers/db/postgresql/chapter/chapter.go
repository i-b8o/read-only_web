package chapter_provider

import (
	"context"

	client "read-only_web/pkg/client/postgresql"

	"read-only_web/internal/domain/entity"
)

type chapterStorage struct {
	client client.PostgreSQLClient
}

func NewChapterStorage(client client.PostgreSQLClient) *chapterStorage {
	return &chapterStorage{client: client}
}

// Get returns an chapter associated with the given ID
func (cs *chapterStorage) Get(ctx context.Context, chapterID uint64) (entity.Chapter, error) {
	const sql = `SELECT id,name,num,order_num,r_id,updated_at FROM "chapter" WHERE id = $1 ORDER BY order_num`
	row := cs.client.QueryRow(ctx, sql, chapterID)
	chapter := entity.Chapter{}
	err := row.Scan(&chapter.ID, &chapter.Name, &chapter.Num, &chapter.OrderNum, &chapter.RegulationID, &chapter.UpdatedAt)
	if err != nil {
		return chapter, err
	}
	return chapter, nil
}

// GetAll returns all chapters associated with the given ID
func (cs *chapterStorage) GetAll(ctx context.Context, regulationID uint64) ([]entity.ChapterInfo, error) {
	const sql = `SELECT id,name,num,order_num FROM "chapter" WHERE r_id = $1 ORDER BY order_num`

	var chapters []entity.ChapterInfo

	rows, err := cs.client.Query(ctx, sql, regulationID)
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

		chapters = append(chapters, chapter)
	}
	return chapters, nil
}
