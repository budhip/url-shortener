package mysql

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/budhip/url-shortener/domain"
)

type mysqlUrlShortenerRepository struct {
	Conn *sql.DB
}

// NewMysqlUrlShortenerRepository will create an object that represent the url_shortener.Repository interface
func NewMysqlUrlShortenerRepository(Conn *sql.DB) domain.UrlShortenerRepository  {
	return &mysqlUrlShortenerRepository{Conn}
}

func (m *mysqlUrlShortenerRepository) GetSlugBySlug (ctx context.Context, shortCode string) (string, error) {
	var slug string

	err := m.Conn.QueryRowContext(ctx,"SELECT slug FROM url_shortener WHERE BINARY slug = ?", shortCode).Scan(&slug)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// there were no rows, but otherwise no error occurred
		} else {
			return "", err
		}
	}

	return slug, nil
}

func (m *mysqlUrlShortenerRepository) Store(ctx context.Context, url, slug string, createdAt time.Time) error {
	query := `INSERT url_shortener SET slug=? , url=? , created_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, errRes := stmt.ExecContext(ctx, slug, url, createdAt)
	if errRes != nil {
		return errRes
	}

	return nil
}
