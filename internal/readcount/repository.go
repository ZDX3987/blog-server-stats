package readcount

import (
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

func (repo *Repository) IsValid(itemId string) bool {
	result, err := repo.db.Query("SELECT COUNT(article_id) valid_count FROM article WHERE article_id = ?", itemId)
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
