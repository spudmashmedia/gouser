package randomuser

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type svc struct {
	host  string
	route string
}

type Service interface {
	GetUsers(count int) (RandomUserResponse, error)
}

func NewService(host string, route string) Service {
	return &svc{
		host:  host,
		route: route,
	}
}

func (s *svc) GetUsers(count int) (RandomUserResponse, error) {
	if count == 0 {
		count = 1 // set default value
	}

	if count > 5000 {
		count = 5000 // cap to max 10
	}

	uri := fmt.Sprintf("%s%s/?results=%d", s.host, s.route, count)

	slog.Debug(
		"randomuser.GetUser: sending request to uri",
		"uri", uri)

	res, err := http.Get(uri)

	if err != nil {
		slog.Error(
			"Randomuser: http client failed",
			"error", err)

		defer res.Body.Close()

		return RandomUserResponse{}, err
	}

	var data RandomUserResponse

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		slog.Error(
			"Randomuser: http response decoding failed",
			"error", err)

		return RandomUserResponse{}, err
	}

	return data, nil
}
