package crawler

import (
	"errors"
	"net/url"
)

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
