package middleware

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"time"
)

type IdempotencyMiddleware struct {
	RedisClient *redis.Client
	TTL         time.Duration
}

func (m *IdempotencyMiddleware) CheckIdempotency(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		idempotencyKey := r.Header.Get("Idempotency-Key")

		if idempotencyKey == "" {
			http.Error(w, "Missing Idempotency-Key header", http.StatusBadRequest)
			return
		}

		// Check if the key exists in Redis
		val, err := m.RedisClient.Get(ctx, idempotencyKey).Result()
		if err == redis.Nil {
			// Key does not exist, proceed and set it
			err = m.RedisClient.Set(ctx, idempotencyKey, "processed", m.TTL).Err()
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		} else if err != nil {
			log.Printf("Redis error: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		} else {
			// Key exists, check the value
			if val == "success" {
				log.Printf("Idempotency key %s already processed successfully", idempotencyKey)
				http.Error(w, "Duplicate request", http.StatusConflict)
				return
			} else if val == "in-progress" {
				log.Printf("Idempotency key %s already in progress", idempotencyKey)
				http.Error(w, "Request is already in progress", http.StatusConflict)
				return
			}
			// Allow retries if the value indicates a failure
			log.Printf("Retrying request for idempotency key %s due to previous failure", idempotencyKey)
		}

		next(w, r)
		// After successful processing, update the key to "success"
		err = m.RedisClient.Set(ctx, idempotencyKey, "success", m.TTL).Err()
		if err != nil {
			log.Printf("Error updating idempotency key status to success: %v", err)
		}
	}
}
