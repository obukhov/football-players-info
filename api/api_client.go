package api

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	TEAM_PATH_TEMPLATE = "/api/teams/en/%d.json"
)

type ApiClientInterface interface {
	GetTeam(id int) (*Team, *ApiError)
}

func NewApiClient() ApiClientInterface {
	endpoint, _ := url.Parse("https://vintagemonster.onefootball.com")

	return &apiClient{
		endpoint: endpoint,
		factory:  buildTeamResponse,
		client:   &httpClient{},
	}
}

type apiClient struct {
	endpoint *url.URL
	factory  teamResponseFactory
	client   httpClientInterface
}

func (ac *apiClient) GetTeam(id int) (*Team, *ApiError) {
	endpoint := *ac.endpoint
	endpoint.Path = fmt.Sprintf(TEAM_PATH_TEMPLATE, id)

	response, err := ac.client.Get(&endpoint)
	if nil != err {
		return nil, &ApiError{
			error:       err,
			recoverable: false,
		}
	}

	if response.StatusCode == http.StatusOK {
		teamResponse, err := ac.factory(response.Body)
		if nil != err {
			return nil, &ApiError{
				error:       err,
				recoverable: true,
			}
		}

		return teamResponse.Data.Team, nil
	}

	recovarable := false
	if response.StatusCode >= http.StatusInternalServerError || response.StatusCode == http.StatusTooManyRequests {
		recovarable = true
	}

	return nil, &ApiError{
		error:       fmt.Errorf("Server returned wrong code %d", response.StatusCode),
		recoverable: recovarable,
	}
}

type ApiError struct {
	error
	recoverable bool
}

func (ae *ApiError) Recoverable() bool {
	return ae.recoverable
}
