package usecase

import (
	"context"
	"regexp"
	"time"

	"github.com/budhip/url-shortener/domain"
	"github.com/budhip/url-shortener/model"
)

type urlShortenerUseCase struct {
	urlShortenerRepo	domain.UrlShortenerRepository
	contextTimeout		time.Duration
}

// NewUrlShortenerUseCase will create new an urlShortenerUseCase object representation of domain.UrlShortenerUseCase interface
func NewUrlShortenerUseCase(usr domain.UrlShortenerRepository, timeout time.Duration) domain.UrlShortenerUseCase {
	return &urlShortenerUseCase{
		urlShortenerRepo:	usr,
		contextTimeout:		timeout,
	}
}

// StoreShorten for checking the shortCode.
// If the shortCode valid and never used then store it into database
func (usu *urlShortenerUseCase) StoreShorten(c context.Context, url, shortCode string) (*model.ShortenResp, error) {
	// check the url not empty
	if url == "" {
		return nil, model.ErrUrlNotPresent
	}

	// Check if the slug exists in the database
	slugFromDB, err := usu.urlShortenerRepo.GetSlugBySlug(c, shortCode)
	if err != nil {
		return nil, model.ErrInternalServerError
	}

	if slugFromDB == shortCode {
		return nil, model.ErrSlugAlreadyUse
	}

	// match between shortCode and regexp
	matched, errMatched := regexp.MatchString(`^[0-9a-zA-Z_]{6}$`, shortCode)
	if errMatched != nil {
		return nil, model.ErrInternalServerError
	}

	if !matched {
		return nil, model.ErrSlugNotMatch
	}

	// store to database
	errStore := usu.urlShortenerRepo.Store(c, url, shortCode, time.Now())
	if errStore != nil {
		return nil, model.ErrInternalServerError
	}

	return &model.ShortenResp{
		Slug: shortCode,
	}, nil
}

// GetShortCode is to get url by shortCode (as parameter)
func (usu *urlShortenerUseCase) GetShortCode(c context.Context, shortCode string) (string, error) {
	data, err := usu.urlShortenerRepo.GetDataBySlug(c, shortCode)
	if err != nil {
		return "", err
	}

	hits := data.Hits + 1

	errUpdate := usu.urlShortenerRepo.Update(c, hits, shortCode, time.Now())
	if errUpdate != nil {
		return "", model.ErrInternalServerError
	}

	return data.URL, nil
}

func (usu *urlShortenerUseCase) GetShortCodeStats(c context.Context, shortCode string) (map[string]interface{}, error) {
	data, err := usu.urlShortenerRepo.GetDataBySlug(c, shortCode)
	if err != nil {
		return nil, err
	}

	if data.Hits == 0 {
		return map[string]interface{}{
			"startDate": data.CreatedAt.UTC().Format(time.RFC3339),
			"redirectCount": data.Hits,
		}, nil
	}

	return map[string]interface{}{
		"startDate": data.CreatedAt.UTC().Format(time.RFC3339),
		"lastSeenDate": data.LastSeen.UTC().Format(time.RFC3339),
		"redirectCount": data.Hits,
	}, nil
}