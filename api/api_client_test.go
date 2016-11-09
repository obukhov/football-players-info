package api

import (
	"github.com/stretchr/testify/mock"
	"net/url"
	"net/http"
	"testing"
)

type httpClientMock struct {
	mock.Mock
}

func (hc *httpClientMock) Get(endpoint *url.URL) (*http.Response, error) {
	args := hc.Called(endpoint)

	return args.Get(0).(*http.Response), args.Error(1)
}

func TestApiClient_GetTeam(t *testing.T) {
	//client := apiClient{}
	// todo
}
