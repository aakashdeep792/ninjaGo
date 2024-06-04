package utils

import (
	"strings"
)

func SkipURL(url string, prefixes []string) bool {
	url = strings.TrimPrefix(url, "/")

	for _, prefix := range prefixes {
		// don't compare when prefix is empty
		if prefix != "" && strings.HasPrefix(url, prefix) {
			return true
		}
	}

	return false
}
