package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/budhip/url-shortener/domain/mocks"
	"github.com/budhip/url-shortener/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	uCase "github.com/budhip/url-shortener/usecase"
)

var (
	mockUrlShortenerRepo = new(mocks.UrlShortenerRepository)
	tempMockShortenResp = model.ShortenResp{}
	url = "https://blog.trello.com/navigate-communication-styles-difficult-times"
	shortCode = "eXaMpL"
)

func TestStoreShorten(t *testing.T) {
	t.Run("url-not-present", func(t *testing.T) {
		url := ""
		shortCut := "eXaMpl"

		u := uCase.NewUrlShortenerUseCase(mockUrlShortenerRepo, time.Second*2)

		_, err := u.StoreShorten(context.TODO(), url, shortCut)

		assert.NotNil(t, err)
		assert.Equal(t, errors.New("url is not present"), err)
		mockUrlShortenerRepo.AssertExpectations(t)
	})

	t.Run("shortcode-already-use", func(t *testing.T) {
		tempMockShortenResp.Slug = "eXaMpl"

		url := "https://blog.trello.com/navigate-communication-styles-difficult-times"
		shortCut := "eXaMpl"

		mockUrlShortenerRepo.On("GetSlugBySlug", mock.Anything, mock.Anything).
			Return(string("eXaMpl"), nil)

		u := uCase.NewUrlShortenerUseCase(mockUrlShortenerRepo, time.Second*2)

		_, err := u.StoreShorten(context.TODO(), url, shortCut)

		assert.NotNil(t, err)
		assert.Equal(t, errors.New("the the desired shortcode is already in use. Shortcodes are case-sensitive"), err)
		mockUrlShortenerRepo.AssertExpectations(t)
	})

	t.Run("shortcode-fails-to-meet-regexp", func(t *testing.T) {
		tempMockShortenResp.Slug = "eXaMpl"

		url := "https://blog.trello.com/navigate-communication-styles-difficult-times"
		shortCut := "eXaMpl65"

		mockUrlShortenerRepo.On("GetSlugBySlug", mock.Anything, mock.Anything).
			Return(string("eXaMpl65"), nil)

		u := uCase.NewUrlShortenerUseCase(mockUrlShortenerRepo, time.Second*2)

		_, err := u.StoreShorten(context.TODO(), url, shortCut)

		assert.NotNil(t, err)
		assert.Equal(t, errors.New("the shortcode fails to meet the following regexp: ^[0-9a-zA-Z_]{6}$"), err)
		mockUrlShortenerRepo.AssertExpectations(t)
	})

	t.Run("store-fails", func(t *testing.T) {
		mockUrlShortenerRepo.On("Store", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(model.ErrInternalServerError)

		u := uCase.NewUrlShortenerUseCase(mockUrlShortenerRepo, time.Second*2)

		_, err := u.StoreShorten(context.TODO(), url, shortCode)
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("internal server error"), err)
		mockUrlShortenerRepo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		mockUrlShortenerRepo := new(mocks.UrlShortenerRepository)

		mockUrlShortenerRepo.On("GetSlugBySlug", mock.Anything, mock.Anything).
			Return("", nil)

		mockUrlShortenerRepo.On("Store", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil)

		u := uCase.NewUrlShortenerUseCase(mockUrlShortenerRepo, time.Second*2)

		res, err := u.StoreShorten(context.TODO(), url, shortCode)
		assert.Nil(t, err)
		assert.Equal(t, shortCode, res.Slug)
		mockUrlShortenerRepo.AssertExpectations(t)
	})

	t.Run("get-slug-fail", func(t *testing.T) {
		mockUrlShortenerRepo := new(mocks.UrlShortenerRepository)

		mockUrlShortenerRepo.On("GetSlugBySlug", mock.Anything, mock.Anything).
			Return("", model.ErrInternalServerError)

		u := uCase.NewUrlShortenerUseCase(mockUrlShortenerRepo, time.Second*2)

		_, err := u.StoreShorten(context.TODO(), url, shortCode)
		assert.NotNil(t, err)
		mockUrlShortenerRepo.AssertExpectations(t)
	})
}

func TestGetShortCode(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUrlShortenerRepo := new(mocks.UrlShortenerRepository)

		mockUrlShortener := model.UrlShortener{Slug: shortCode, URL: url, Hits: 0, CreatedAt: time.Now(), LastSeen: time.Now()}

		mockUrlShortenerRepo.On("GetDataBySlug", mock.Anything, mock.Anything).
			Return(mockUrlShortener, nil)

		mockUrlShortenerRepo.On("Update", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil)

		u := uCase.NewUrlShortenerUseCase(mockUrlShortenerRepo, time.Second*2)

		res, err := u.GetShortCode(context.TODO(), shortCode)
		assert.Nil(t, err)
		assert.Equal(t, mockUrlShortener.URL, res)
		mockUrlShortenerRepo.AssertExpectations(t)
	})

	t.Run("get-data-failed", func(t *testing.T) {
		mockUrlShortenerRepo.On("GetDataBySlug", mock.Anything, mock.Anything).
			Return(model.UrlShortener{}, model.ErrInternalServerError)

		u := uCase.NewUrlShortenerUseCase(mockUrlShortenerRepo, time.Second*2)

		_, err := u.GetShortCode(context.TODO(), shortCode)
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("internal server error"), err)
		mockUrlShortenerRepo.AssertExpectations(t)
	})
}

func TestGetShortCodeStats(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUrlShortenerRepo := new(mocks.UrlShortenerRepository)

		mockUrlShortener := model.UrlShortener{Slug: shortCode, URL: url, Hits: 0, CreatedAt: time.Now(), LastSeen: time.Now()}

		mockUrlShortenerRepo.On("GetDataBySlug", mock.Anything, mock.Anything).
			Return(mockUrlShortener, nil)

		u := uCase.NewUrlShortenerUseCase(mockUrlShortenerRepo, time.Second*2)

		res, err := u.GetShortCodeStats(context.TODO(), shortCode)
		assert.Nil(t, err)
		assert.Equal(t, mockUrlShortener.Hits, res["redirectCount"])
		mockUrlShortenerRepo.AssertExpectations(t)
	})

	t.Run("success-with-hits", func(t *testing.T) {
		mockUrlShortenerRepo := new(mocks.UrlShortenerRepository)

		mockUrlShortener := model.UrlShortener{Slug: shortCode, URL: url, Hits: 1, CreatedAt: time.Now(), LastSeen: time.Now()}

		mockUrlShortenerRepo.On("GetDataBySlug", mock.Anything, mock.Anything).
			Return(mockUrlShortener, nil)

		u := uCase.NewUrlShortenerUseCase(mockUrlShortenerRepo, time.Second*2)

		res, err := u.GetShortCodeStats(context.TODO(), shortCode)
		assert.Nil(t, err)
		assert.Equal(t, mockUrlShortener.Hits, res["redirectCount"])
		mockUrlShortenerRepo.AssertExpectations(t)
	})

	t.Run("success-with-hits", func(t *testing.T) {
		mockUrlShortenerRepo.On("GetDataBySlug", mock.Anything, mock.Anything).
			Return(nil, model.ErrInternalServerError)

		u := uCase.NewUrlShortenerUseCase(mockUrlShortenerRepo, time.Second*2)

		_, err := u.GetShortCodeStats(context.TODO(), shortCode)
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("internal server error"), err)
		mockUrlShortenerRepo.AssertExpectations(t)
	})
}
