package pageview

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) SavePageView(ctx context.Context, mpv []*PageView) (bool, error) {
	sqlASegment := make([]string, len(mpv))
	args := make([]any, 0)
	for i, value := range mpv {
		sqlASegment[i] = `(?, ?, ?, ?, ?, ?, ?, ?, ?)`
		args = append(args, value.EventID, value.VisitorID, value.Path, value.Referer, value.UserAgent, value.IP,
			value.DurationSec, value.MaxScrollDepth, value.OccurredAt)
	}
	sqlStr := fmt.Sprintf(`
INSERT INTO traffic_page_view(event_id, visitor_id, path, referer, user_agent, ip, duration_sec, max_scroll_depth, occurred_at) 
VALUES %s;
`, strings.Join(sqlASegment, ", "))
	result, err := repo.db.ExecContext(ctx, sqlStr, args...)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	return affected > 0, err
}
