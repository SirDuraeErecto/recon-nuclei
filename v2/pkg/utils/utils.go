package utils

import (
	"errors"
	"net/url"
	"strings"

	"github.com/projectdiscovery/fileutil"
)

func IsBlank(value string) bool {
	return strings.TrimSpace(value) == ""
}

func IsNotBlank(value string) bool {
	return !IsBlank(value)
}

func UnwrapError(err error) error {
	for { // get the last wrapped error
		unwrapped := errors.Unwrap(err)
		if unwrapped == nil {
			break
		}
		err = unwrapped
	}
	return err
}

func LoadFile(filename string) ([]string, error) {
	var items []string
	readfileChan, err := fileutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	for includeIdLine := range readfileChan {
		items = append(items, includeIdLine)
	}
	return items, nil
}

// IsURL tests a string to determine if it is a well-structured url or not.
func IsURL(input string) bool {
	_, err := url.ParseRequestURI(input)
	if err != nil {
		return false
	}

	u, err := url.Parse(input)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
