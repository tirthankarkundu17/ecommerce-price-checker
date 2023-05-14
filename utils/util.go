package utils

import (
	"net/url"
	"strings"
)

func RemoveQueryParamsFromUrl(url string) string {
	return Sanitize(strings.Split(url, "?")[0])
}

func Sanitize(s string) string {
	// Replace < and > with their HTML entity equivalents
	s = strings.Replace(s, "<", "&lt;", -1)
	s = strings.Replace(s, ">", "&gt;", -1)
	return s
}

func GetDomainHost(urlString string) (string, error) {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}

	return parsedURL.Hostname(), nil
}
