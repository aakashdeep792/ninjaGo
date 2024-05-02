package utils

import "strings"

func SkipURL(url string, prefixes []string) bool {
	url = strings.TrimPrefix(url, "/")

	for _, prefix := range prefixes {
		if strings.HasPrefix(url, prefix) {
			return true
		}
	}

	return false
}
