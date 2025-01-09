package utils

import (
	"context"
	urls_mapping "github.com/Rammurthy5/url-shortner-go/internal/db/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

type MockDBTX struct {
	mock.Mock
}

func (m *MockDBTX) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	arguments := m.Called(ctx, query, args)
	return arguments.Get(0).(pgconn.CommandTag), arguments.Error(1)
}

func (m *MockDBTX) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	arguments := m.Called(ctx, query, args)
	return arguments.Get(0).(pgx.Rows), arguments.Error(1)
}

func (m *MockDBTX) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	arguments := m.Called(ctx, query, args)
	return arguments.Get(0).(pgx.Row)
}

func TestFetchShortURL(t *testing.T) {
	mockDB := new(MockDBTX)
	queries := urls_mapping.New(mockDB)

	url := "https://example.com"
	expectedShortUrl := "e2fc714c"

	// Mock QueryRow behavior
	mockDB.On("QueryRow", mock.Anything, mock.Anything, url).
		Return(&urls_mapping.UrlMapping{ShortUrl: expectedShortUrl}, nil)

	result := FetchShortURL(queries, url)
	require.Equal(t, expectedShortUrl, result)

	// Simulate error
	mockDB.On("QueryRow", mock.Anything, mock.Anything, url).
		Return(nil, errors.New("not found"))

	result = FetchShortURL(queries, url)
	require.Equal(t, "", result)
}

func TestStoreShortURL(t *testing.T) {
	mockDB := new(MockDBTX)
	queries := urls_mapping.New(mockDB)

	longUrl := "https://example.com"
	shortUrl := "e2fc714c"

	mockDB.On("QueryRow", mock.Anything, mock.Anything, longUrl, shortUrl).
		Return(&urls_mapping.UrlMapping{}, nil)

	err := StoreShortURL(queries, longUrl, shortUrl)
	require.NoError(t, err)

	// Simulate error
	mockDB.On("QueryRow", mock.Anything, mock.Anything, longUrl, shortUrl).
		Return(nil, errors.New("insert error"))

	err = StoreShortURL(queries, longUrl, shortUrl)
	require.Error(t, err)
}

func TestDeleteShortURL(t *testing.T) {
	mockDB := new(MockDBTX)
	queries := urls_mapping.New(mockDB)

	longUrl := "https://example.com"

	mockDB.On("Exec", mock.Anything, mock.Anything, longUrl).
		Return(pgconn.CommandTag{}, nil)

	err := DeleteShortURL(queries, longUrl)
	require.NoError(t, err)

	// Simulate error
	mockDB.On("Exec", mock.Anything, mock.Anything, longUrl).
		Return(pgconn.CommandTag{}, errors.New("delete error"))

	err = DeleteShortURL(queries, longUrl)
	require.Error(t, err)
}

func TestUpdateShortURL(t *testing.T) {
	mockDB := new(MockDBTX)
	queries := urls_mapping.New(mockDB)

	longUrl := "https://example.com"
	shortUrl := "e2fc714c"

	mockDB.On("QueryRow", mock.Anything, mock.Anything, shortUrl, longUrl).
		Return(&urls_mapping.UrlMapping{}, nil)

	err := UpdateShortURL(queries, longUrl, shortUrl)
	require.NoError(t, err)

	// Simulate error
	mockDB.On("QueryRow", mock.Anything, mock.Anything, shortUrl, longUrl).
		Return(nil, errors.New("update error"))

	err = UpdateShortURL(queries, longUrl, shortUrl)
	require.Error(t, err)
}
