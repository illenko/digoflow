package task

import (
	"fmt"
	"strings"

	"github.com/illenko/digoflow/container"

	"github.com/go-resty/resty/v2"
)

var HttpClient = resty.New().SetDebug(true)

// HTTPRequest is a task that sends an HTTP request
func HTTPRequest(_ *container.Container, input Input) (Output, error) {
	requestConfig, err := parseInput(input)
	if err != nil {
		return nil, err
	}

	response, err := executeRequest(requestConfig)
	if err != nil {
		return nil, err
	}

	return response, nil
}

const (
	HeaderPrefix = "header."
	QueryPrefix  = "query."
	BodyPrefix   = "body."
)

// httpRequestConfig is a struct that holds the configuration for an HTTP request
type httpRequestConfig struct {
	uri         string
	method      string
	headers     map[string]string
	queryParams map[string]string
	body        map[string]any
}

// parseInput parses the input for the HTTP request task
// The input should contain the following
// uri: string (required)
// method: string (required) (GET, POST, PUT, DELETE, PATCH)
// headers: map[string]string (optional)
// query parameters: map[string]string (optional)
// body: map[string]any (optional)
func parseInput(input Input) (httpRequestConfig, error) {
	uri, ok := input["uri"].(string)
	if !ok {
		return httpRequestConfig{}, fmt.Errorf("uri not found or not a string")
	}

	method, ok := input["method"].(string)
	if !ok {
		return httpRequestConfig{}, fmt.Errorf("method not found or not a string")
	}

	headers := map[string]string{}
	queryParams := map[string]string{}
	body := map[string]any{}

	for k, v := range input {
		switch {
		case strings.HasPrefix(k, HeaderPrefix):
			headers[strings.TrimPrefix(k, HeaderPrefix)] = v.(string)
		case strings.HasPrefix(k, QueryPrefix):
			queryParams[strings.TrimPrefix(k, QueryPrefix)] = v.(string)
		case strings.HasPrefix(k, BodyPrefix):
			body[strings.TrimPrefix(k, BodyPrefix)] = v
		}
	}

	return httpRequestConfig{
		uri:         uri,
		method:      method,
		headers:     headers,
		queryParams: queryParams,
		body:        body,
	}, nil
}

// executeRequest executes the HTTP request
func executeRequest(config httpRequestConfig) (Output, error) {
	response := map[string]any{}
	errorResponse := map[string]any{}

	resp, err := HttpClient.R().
		SetHeaders(config.headers).
		SetQueryParams(config.queryParams).
		SetBody(config.body).
		SetResult(&response).
		SetError(&errorResponse).
		Execute(config.method, config.uri)

	if err != nil {
		return nil, err
	}

	output := Output{}

	output["http-status"] = resp.Status()
	output["http-status-code"] = resp.StatusCode()
	output["success"] = !resp.IsError()

	if resp.IsError() {
		for k, v := range errorResponse {
			output[k] = v
		}
	} else {
		for k, v := range response {
			output[k] = v
		}
	}
	return output, nil
}
