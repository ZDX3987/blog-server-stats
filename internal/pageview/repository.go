package pageview

import (
	"context"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func (repo *Repository) SavePageView(ctx context.Context, pv *PageView) (bool, error) {
	sqlStr := `
INSERT INTO traffic_page_view(event_id, visitor_id, path, referer, user_agent, ip, duration_sec, max_scroll_depth) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
`
	result, err := repo.db.ExecContext(ctx, sqlStr)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	return affected > 0, err
}
