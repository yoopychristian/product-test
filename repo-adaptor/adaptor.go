package repositoryadaptor

import (
	"encoding/json"
	"fmt"
	"net/url"

	fx "product-test/functions"

	"go.uber.org/zap"
)

type RepositoryURL struct {
	Read  string `json:"read_url"`
	Write string `json:"write_url"`
}
type HTTPResponse struct {
	Status      bool   `json:"status"`
	ErrorCode   string `json:"error_code"`
	Description string `json:"description"`
	Data        string `json:"data"`
}

type RepositoryAdaptor struct {
	URL        RepositoryURL
	Client     HttpClient
	ServiceURL string
	Master     RepositoryURL
}

func (adaptor RepositoryAdaptor) GET(l *zap.Logger, uri string, data interface{}) (HTTPResponse, error) {
	handleErr := func(err error) (HTTPResponse, error) {
		return HTTPResponse{}, fmt.Errorf("get %s  : %w", uri, err)
	}

	baseUrl, err := url.Parse(uri)
	if err != nil {
		return handleErr(err)
	}

	// Add a Path Segment (Path segment is automatically escaped)
	params, err := fx.StructToUrlValue(data)
	if err != nil {
		return handleErr(err)
	}

	// Add Query Parameters to the URL
	baseUrl.RawQuery = params.Encode() // Escape Query Parameters

	l.Debug("http request",
		zap.String("method", "GET"),
		zap.String("url", baseUrl.String()))

	result, err := adaptor.Client.GET(baseUrl.String())
	if err != nil {
		return handleErr(fmt.Errorf("http process (%w)", err))
	}

	rr := HTTPResponse{}
	if err := json.Unmarshal(result, &rr); err != nil {
		return handleErr(fmt.Errorf("unmarshal response , %s (%w)", string(result), err))
	}

	return rr, nil
}

func (adaptor RepositoryAdaptor) POST(l *zap.Logger, url string, data interface{}) (HTTPResponse, error) {
	handleErr := func(err error) (HTTPResponse, error) {
		return HTTPResponse{}, fmt.Errorf("post %s  : %w", url, err)
	}

	message, err := json.Marshal(data)
	if err != nil {
		return handleErr(err)
	}

	l.Debug("http request",
		zap.String("method", "POST"),
		zap.String("ur", url),
		zap.Any("data", data))

	result, err := adaptor.Client.POST(url, message)
	if err != nil {
		return handleErr(fmt.Errorf("http process (%w)", err))
	}

	rr := HTTPResponse{}
	if err := json.Unmarshal(result, &rr); err != nil {
		return handleErr(fmt.Errorf("unmarshal response , %s (%w)", string(result), err))
	}

	return rr, nil
}

func (adaptor RepositoryAdaptor) PUT(l *zap.Logger, url string, data interface{}) (HTTPResponse, error) {
	handleErr := func(err error) (HTTPResponse, error) {
		return HTTPResponse{}, fmt.Errorf("put %s  : %w", url, err)
	}

	message, err := json.Marshal(data)
	if err != nil {
		return handleErr(err)
	}

	l.Debug("http request",
		zap.String("method", "PUT"),
		zap.String("ur", url),
		zap.Any("data", data))

	result, err := adaptor.Client.PUT(url, message)
	if err != nil {
		return handleErr(fmt.Errorf("http process (%w)", err))
	}

	rr := HTTPResponse{}
	if err := json.Unmarshal(result, &rr); err != nil {
		return handleErr(fmt.Errorf("unmarshal response , %s (%w)", string(result), err))
	}

	return rr, nil
}

func (adaptor RepositoryAdaptor) DELETE(l *zap.Logger, url string, data interface{}) (HTTPResponse, error) {
	handleErr := func(err error) (HTTPResponse, error) {
		return HTTPResponse{}, fmt.Errorf("delete %s  : %w", url, err)
	}

	message, err := json.Marshal(data)
	if err != nil {
		return handleErr(err)
	}

	l.Debug("http request",
		zap.String("method", "DELETE"),
		zap.String("ur", url),
		zap.Any("data", data))

	result, err := adaptor.Client.DELETE(url, message)
	if err != nil {
		return handleErr(fmt.Errorf("http process (%w)", err))
	}

	rr := HTTPResponse{}
	if err := json.Unmarshal(result, &rr); err != nil {
		return handleErr(fmt.Errorf("unmarshal response , %s (%w)", string(result), err))
	}

	return rr, nil
}
