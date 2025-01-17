package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
)

func TestCheckIdempotency(t *testing.T) {
	// Create a mock Redis client
	mockRedis, mock := redismock.NewClientMock()
	middleware := &IdempotencyMiddleware{
		RedisClient: mockRedis,
		TTL:         5 * time.Second,
	}

	// Define a sample handler to wrap with middleware
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		// nolint:errcheck
		w.Write([]byte("Success"))
	}

	t.Run("Missing Idempotency-Key", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/shorten", nil)
		rec := httptest.NewRecorder()

		wrappedHandler := middleware.CheckIdempotency(handler)
		wrappedHandler(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "Missing Idempotency-Key")
	})

	t.Run("New Idempotency-Key", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/shorten", nil)
		req.Header.Set("Idempotency-Key", "unique-key-1")
		rec := httptest.NewRecorder()

		// Mock Redis: Key does not exist
		mock.ExpectGet("unique-key-1").RedisNil()
		mock.ExpectSet("unique-key-1", "in-progress", 5*time.Second).SetVal("OK")

		wrappedHandler := middleware.CheckIdempotency(handler)
		wrappedHandler(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Success", rec.Body.String())
		mock.ClearExpect()
	})

	t.Run("Duplicate Idempotency-Key Success", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/shorten", nil)
		req.Header.Set("Idempotency-Key", "unique-key-2")
		rec := httptest.NewRecorder()

		// Mock Redis: Key already exists with "success"
		mock.ExpectGet("unique-key-2").SetVal("success")

		wrappedHandler := middleware.CheckIdempotency(handler)
		wrappedHandler(rec, req)

		assert.Equal(t, http.StatusConflict, rec.Code)
		assert.Contains(t, rec.Body.String(), "Duplicate request")
		mock.ClearExpect()
	})

	t.Run("Retry After Failure", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/shorten", nil)
		req.Header.Set("Idempotency-Key", "unique-key-3")
		rec := httptest.NewRecorder()

		// Mock Redis: Key exists with "failure"
		mock.ExpectGet("unique-key-3").SetVal("failure")
		mock.ExpectSet("unique-key-3", "in-progress", 24*time.Hour).SetVal("OK")

		wrappedHandler := middleware.CheckIdempotency(handler)
		wrappedHandler(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Success", rec.Body.String())
		mock.ClearExpect()
	})

	t.Run("Redis Error Handling", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/shorten", nil)
		req.Header.Set("Idempotency-Key", "unique-key-4")
		rec := httptest.NewRecorder()

		// Mock Redis: Redis operation error
		mock.ExpectGet("unique-key-4").SetErr(context.DeadlineExceeded)

		wrappedHandler := middleware.CheckIdempotency(handler)
		wrappedHandler(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), "Internal Server Error")
		mock.ClearExpect()
	})
}
