package utils

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	urls_mapping "github.com/Rammurthy5/url-shortner-go/internal/db/sqlc"
)

func Shorten(url string) string {
	h := sha256.New()
	h.Write([]byte(url))
	hashed := hex.EncodeToString(h.Sum(nil))
	return hashed[:8]
}

func FetchShortURL(dbInst *urls_mapping.Queries, url string) string {
	shortUrl, err := dbInst.GetUrl(context.Background(), url)
	if err != nil {
		fmt.Sprintf("error while fetching short url: %v", err)
		return ""
	}
	return shortUrl.ShortUrl
}

func StoreShortURL(dbInst *urls_mapping.Queries, longUrl string, shortUrl string) error {
	_, err := dbInst.InsertUrl(context.Background(), urls_mapping.InsertUrlParams{ShortUrl: shortUrl, LongUrl: longUrl})
	if err != nil {
		return err
	}
	return nil
}

func DeleteShortURL(dbInst *urls_mapping.Queries, longUrl string) error {
	err := dbInst.DeleteUrl(context.Background(), longUrl)
	if err != nil {
		return err
	}
	return nil
}

func UpdateShortURL(dbInst *urls_mapping.Queries, longUrl string, shortUrl string) error {
	_, err := dbInst.UpdateUrl(context.Background(), urls_mapping.UpdateUrlParams{ShortUrl: shortUrl, LongUrl: longUrl})
	if err != nil {
		return err
	}
	return nil
}
