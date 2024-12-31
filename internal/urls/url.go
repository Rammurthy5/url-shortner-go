package urls

import (
	"crypto/sha256"
	"encoding/hex"
)

func Shorten(url string) string {
	h := sha256.New()
	h.Write([]byte(url))
	hashed := hex.EncodeToString(h.Sum(nil))
	return hashed[:8]
}
