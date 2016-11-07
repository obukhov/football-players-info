package api

import (
	"net/url"
	"net/http"
)

type httpClientInterface interface {
	Get(*url.URL) (*http.Response, error)
}

type httpClient struct {
}

func (hc *httpClient) Get(endpoint *url.URL) (*http.Response, error) {
	request, err := http.NewRequest("GET", endpoint.String(), nil)
	if nil != err {
		return nil, err
	}

	client := http.Client{}
	response, err := client.Do(request)

	return response, err
}
