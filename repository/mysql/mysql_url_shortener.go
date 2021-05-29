package mysql

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/budhip/url-shortener/domain"
	"github.com/budhip/url-shortener/model"
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

func (m *mysqlUrlShortenerRepository) GetDataBySlug (ctx context.Context, shortCode string) (model.UrlShortener, error) {
	us := model.UrlShortener{}
	query := `SELECT slug, url, created_at, last_seen, hits FROM url_shortener WHERE BINARY slug = ?`

	err := m.Conn.QueryRowContext(ctx, query, shortCode).Scan(
		&us.Slug,
		&us.URL,
		&us.CreatedAt,
		&us.LastSeen,
		&us.Hits,
	)
	if err != nil {
		log.Println("err GetDataBySlug repo: ", err)
		if errors.Is(err, sql.ErrNoRows) {
			return model.UrlShortener{}, model.ErrSlugNotFound
		} else {
			return model.UrlShortener{}, err
		}
	}

	return us, nil
}

func (m *mysqlUrlShortenerRepository) Update (ctx context.Context, hits int, shortCode string, lastSeen time.Time) error {
	query := `UPDATE url_shortener SET last_seen=?, hits=? WHERE BINARY slug = ?`

	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		log.Println("err update: ", err)
		return model.ErrInternalServerError
	}
	defer stmt.Close()

	_, errExec := stmt.ExecContext(ctx, lastSeen, hits, shortCode)
	if errExec != nil {
		return model.ErrInternalServerError
	}

	return nil
}