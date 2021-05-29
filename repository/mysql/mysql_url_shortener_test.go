package mysql_test

import (
	"context"
	"testing"
	"time"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/stretchr/testify/assert"

	urlShortMysqlRepo "github.com/budhip/url-shortener/repository/mysql"
)

func TestStore(t *testing.T) {
	now := time.Now()
	url := "https://blog.trello.com/navigate-communication-styles-difficult-times"
	shortCode := "eXaMpl"


	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "INSERT url_shortener SET slug=\\? , url=\\? , created_at=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(shortCode, url, now).
		WillReturnResult(sqlmock.NewResult(1, 5))

	urlShort:= urlShortMysqlRepo.NewMysqlUrlShortenerRepository(db)

	err = urlShort.Store(context.TODO(), url, shortCode, now)
	assert.NoError(t, err)
}

func TestStoreQueryFailed(t *testing.T) {
	now := time.Now()
	url := "https://blog.trello.com/navigate-communication-styles-difficult-times"
	shortCode := "eXaMpl"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "INSERT url_shortener SET slug= \\? , url=\\? , created_at=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(shortCode, url, now).
		WillReturnResult(sqlmock.NewResult(1, 5))

	urlShort:= urlShortMysqlRepo.NewMysqlUrlShortenerRepository(db)

	err = urlShort.Store(context.TODO(), url, shortCode, now)
	assert.Error(t, err)
	assert.NotNil(t, err)
}

func TestStoreFailed(t *testing.T) {
	now := time.Now()
	url := "https://blog.trello.com/navigate-communication-styles-difficult-times"
	shortCode := "eXaMpl"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "INSERT url_shortener SET slug=\\? , url=\\? , created_at=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(shortCode, url, "1").
		WillReturnResult(sqlmock.NewResult(1, 5))

	urlShort:= urlShortMysqlRepo.NewMysqlUrlShortenerRepository(db)

	err = urlShort.Store(context.TODO(), url, shortCode, now)
	assert.Error(t, err)
	assert.NotNil(t, err)
}
