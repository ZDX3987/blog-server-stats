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
	result, err := repo.db.Query("select count(article_id) valid_count from article where article_id = ?", itemId)
	if err != nil {
		return false
	}
	defer result.Close()
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
