package domain

import (
	"context"
	"time"

	"github.com/budhip/url-shortener/model"
)

// UrlShortenerUseCase represent the url_shortener's use cases
type UrlShortenerUseCase interface {
	StoreShorten(ctx context.Context, url, shortCode string) (*model.ShortenResp, error)
}

// UrlShortenerRepository represent the url_shortener's repository contract
type UrlShortenerRepository interface {
	GetSlugBySlug(ctx context.Context, shortCode string) (string, error)
	Store(ctx context.Context, url, slug string, createdAt time.Time) error
}
