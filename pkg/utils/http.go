package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

type HttpRequest struct {
	http.Client
	Response *http.Response
	Error    error
}

func (hr *HttpRequest) Get(url string, headers map[string]string) *HttpRequest {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		hr.Error = err
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	hr.Response, hr.Error = hr.Do(req)

	return hr
}

func (hr *HttpRequest) Post(url string, headers map[string]string, body io.Reader) *HttpRequest {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		hr.Error = err
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	hr.Response, hr.Error = hr.Do(req)

	return hr
}

// ParseJson Parse the return value into json format
func (hr *HttpRequest) ParseJson(payload any) error {
	bytes, err := hr.ParseBytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, &payload)
}

// ParseBytes Parse the return value into []byte format
func (hr *HttpRequest) ParseBytes() ([]byte, error) {
	if hr.Error != nil {
		return nil, hr.Error
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err.Error())
		}
	}(hr.Response.Body)

	return io.ReadAll(hr.Response.Body)
}

// Raw Return the raw response data as a string
func (hr *HttpRequest) Raw() (string, error) {
	str, err := hr.ParseBytes()
	if err != nil {
		return "", err
	}
	return string(str), nil
}
