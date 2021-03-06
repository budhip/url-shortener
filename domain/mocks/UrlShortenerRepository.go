// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/budhip/url-shortener/model"

	time "time"
)

// UrlShortenerRepository is an autogenerated mock type for the UrlShortenerRepository type
type UrlShortenerRepository struct {
	mock.Mock
}

// GetDataBySlug provides a mock function with given fields: ctx, shortCode
func (_m *UrlShortenerRepository) GetDataBySlug(ctx context.Context, shortCode string) (model.UrlShortener, error) {
	ret := _m.Called(ctx, shortCode)

	var r0 model.UrlShortener
	if rf, ok := ret.Get(0).(func(context.Context, string) model.UrlShortener); ok {
		r0 = rf(ctx, shortCode)
	} else {
		r0 = ret.Get(0).(model.UrlShortener)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, shortCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSlugBySlug provides a mock function with given fields: ctx, shortCode
func (_m *UrlShortenerRepository) GetSlugBySlug(ctx context.Context, shortCode string) (string, error) {
	ret := _m.Called(ctx, shortCode)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, shortCode)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, shortCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, url, slug, createdAt
func (_m *UrlShortenerRepository) Store(ctx context.Context, url string, slug string, createdAt time.Time) error {
	ret := _m.Called(ctx, url, slug, createdAt)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, time.Time) error); ok {
		r0 = rf(ctx, url, slug, createdAt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, hits, shortCode, lastSeen
func (_m *UrlShortenerRepository) Update(ctx context.Context, hits int, shortCode string, lastSeen time.Time) error {
	ret := _m.Called(ctx, hits, shortCode, lastSeen)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string, time.Time) error); ok {
		r0 = rf(ctx, hits, shortCode, lastSeen)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
