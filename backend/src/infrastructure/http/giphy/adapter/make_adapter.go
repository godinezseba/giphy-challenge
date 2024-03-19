package adapter

import "net/http"

type gifHTTPAdapter struct {
	httpClient *http.Client
}

func MakeGIFAdapter(
	httpClient *http.Client,
) *gifHTTPAdapter {
	return &gifHTTPAdapter{
		httpClient: httpClient,
	}
}
