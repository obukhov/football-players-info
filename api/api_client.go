package api

import (
	"fmt"
	"net/url"
)

const (
	TEAM_PATH_TEMPLATE = "/api/teams/en/%d.json"
)

type ApiClientInterface interface {
	GetTeam(id int) (*Team, error)
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

func (ac *apiClient) GetTeam(id int) (*Team, error) {
	endpoint := *ac.endpoint
	endpoint.Path = fmt.Sprintf(TEAM_PATH_TEMPLATE, id)

	response, err := ac.client.Get(&endpoint)
	if nil != err {
		return nil, err
	}

	teamResponse, err := ac.factory(response.Body)
	if nil != err {
		return nil, err
	}

	return teamResponse.Data.Team, nil
}
