package api

import (
	"io"
	"encoding/json"
)

type teamResponseFactory func(body io.ReadCloser) (*TeamResponse, error)

func buildTeamResponse(body io.ReadCloser) (*TeamResponse, error) {
	decoder := json.NewDecoder(body)
	result := &TeamResponse{}
	result.Data.Team = &Team{}

	if err := decoder.Decode(result); nil != err {
		return nil, err
	}

	return result, nil
}
