package repositories

import (
	"context"
	"github.com/ilmedova/go-url-shortener/internal/models"
	"github.com/jmoiron/sqlx"
)

type URLRepository struct {
	DB *sqlx.DB
}

func NewURLRepository(db *sqlx.DB) *URLRepository {
	return &URLRepository{DB: db}
}

func (r *URLRepository) SaveURL(ctx context.Context, url models.URL) error {
	query := `INSERT INTO urls (id, original, shortened) VALUES ($1, $2, $3)`
	_, err := r.DB.ExecContext(ctx, query, url.ID, url.Original, url.Shortened)
	return err
}

func (r *URLRepository) GetOriginalURL(ctx context.Context, short string) (*models.URL, error) {
	var url models.URL
	query := `SELECT * FROM urls WHERE shortened = $1`
	err := r.DB.GetContext(ctx, &url, query, short)
	return &url, err
}
