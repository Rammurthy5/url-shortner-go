package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShorten(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want string
	}{
		{name: "Basic URL", url: "https://trumpwhitehouse.archives.gov//bill-122720/", want: "[0-9a-f]{8}"},
		{name: "Empty URL", url: "", want: "[0-9a-f]{8}"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Shorten(tt.url)
			assert.Regexp(t, tt.want, got)
		})
	}
}
