package client

import (
	"net/url"
	"path/filepath"
)

func AppendToURL(baseURL *url.URL, endpointSegments ...string) *url.URL {
	// make a COPY of the URL, so that when we add to the path, we don't change
	// the value of baseURL.
	reqURL := *baseURL

	pathParts := append([]string{reqURL.Path}, endpointSegments...)
	reqURL.Path = filepath.Join(pathParts...)
	return &reqURL
}

// AppendQueryToURL will append the query params and return the updated query
func AppendQueryToURL(reqURL *url.URL, params map[string]string) *url.URL {
	queryParams := url.Values{}
	for key, value := range params {
		queryParams.Add(key, value)
	}
	reqURL.RawQuery = queryParams.Encode()
	return reqURL
}