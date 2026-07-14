package readcount

import (
	"context"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db}
}

type validItemCount struct {
	count int
}

func (repo *Repository) IsValid(ctx context.Context, itemId string) bool {
	result, err := repo.db.QueryContext(ctx, "SELECT COUNT(article_id) valid_count FROM article WHERE article_id = ?", itemId)
	if err != nil {
		return false
	}
	defer func(result *sql.Rows) {
		err := result.Close()
		if err != nil {
			return
		}
	}(result)
	var validCount validItemCount

	for result.Next() {
		err := result.Scan(&validCount.count)
		if err != nil {
			return false
		}
	}

	if err := result.Err(); err != nil {
		return false
	}
	return validCount.count > 0
}

func (repo *Repository) InsertReadLog(ctx context.Context, log *ReadCountLog) (bool, error) {
	query := `
INSERT INTO read_count_log(item_id, item_type, identity, visitor_id, visitor_ip, user_agent, referer, read_duration, read_depth)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, )
`
	result, err := repo.db.ExecContext(ctx, query, log.ItemId, 1, log.Identity, log.VisitorID, log.VisitorIp, log.UserAgent, log.Referer, log.ReadDuration, log.ReadDepth)
	if err != nil {
		return false, err
	}
	rowNum, err := result.RowsAffected()
	return rowNum > 0, err
}

func (repo *Repository) UpdateReadCount(ctx context.Context, itemID string, count int64) error {
	if count == 0 {
		return nil
	}
	s := "UPDATE article SET read_count = read_count + ? WHERE article_id = ?"
	_, err := repo.db.ExecContext(ctx, s, count, itemID)
	return err
}
