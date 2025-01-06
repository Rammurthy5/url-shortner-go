package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestShorten(t *testing.T) {
	url := "https://trumpwhitehouse.archives.gov/briefings-statements/bill-announcement-122720/"

	shortenedURL := Shorten(url)
	require.Equal(t, 8, len(shortenedURL))
	require.NotEqual(t, url, shortenedURL)
}
