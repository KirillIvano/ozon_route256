package jsonreqwrap

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type HttpJsonClient[Req, Res any] struct {
	Url    string
	Method string
}

func NewClient[Req, Res any](url, method string) HttpJsonClient[Req, Res] {
	return HttpJsonClient[Req, Res]{
		Url:    url,
		Method: method,
	}
}

func (client *HttpJsonClient[Req, Res]) Run(body Req) (*Res, error) {
	rawJson, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrap(err, "encoding to json")
	}

	req, err := http.NewRequest(client.Method, client.Url, bytes.NewBuffer(rawJson))
	if err != nil {
		return nil, errors.Wrap(err, "creating request")
	}

	httpResponse, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "initializing request")
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusOK {
		return nil, errors.Errorf("http request error: %s", httpResponse.Status)
	}

	var response Res
	err = json.NewDecoder(httpResponse.Body).Decode(&response)

	if err != nil {
		return nil, errors.Wrap(err, "json decode error")
	}

	return &response, nil
}
