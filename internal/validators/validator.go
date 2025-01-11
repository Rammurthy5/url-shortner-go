package validators

import (
	"fmt"
	"regexp"
	"strings"
)

// validateURL checks if the URL is valid.
func ValidateURL(url string) error {
	// Check if the URL is empty
	if url == "" {
		return fmt.Errorf("URL cannot be empty")
	}

	// Check if the URL starts with http:// or https://
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return fmt.Errorf("URL must start with http:// or https://")
	}

	// Optional: You can use a more robust regex to validate the URL structure
	// This is a very basic validation for demonstration.
	re := `^(http|https)://([a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)+)(:[0-9]+)?(/.*)?$`
	r := regexp.MustCompile(re)
	if !r.MatchString(url) {
		return fmt.Errorf("Invalid URL format")
	}

	return nil
}
