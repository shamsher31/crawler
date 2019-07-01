package crawler

import (
	"errors"
	"net/url"
)

// Validate accepts URL and check if it is valid
func Validate(URL string) error {
	if URL == "" {
		return errors.New("URL argument is missing")
	}
	// Verify if the URL is valid
	_, err := url.ParseRequestURI(URL)
	if err != nil {
		return errors.New("Invalid URL")
	}
	return nil
}
