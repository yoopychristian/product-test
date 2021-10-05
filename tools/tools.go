package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
	"go.uber.org/zap"
)

type RepositoryAdaptor struct {
	Client             HttpClient
	ReadRepositoryURL  string
	WriteRepositoryURL string
	PortalServiceURL   string
	BOServiceURL       string
	MasterServiceURL   string
	GatewayURL         string
}

type HTTPResponse struct {
	Status      bool   `json:"status"`
	ErrorCode   string `json:"error_code"`
	Description string `json:"description"`
	Data        string `json:"data"`
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
	params, err := StructToUrlValue(data)
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
		zap.String("url", url),
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
		zap.String("url", url),
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
		zap.String("url", url),
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

type RepositoryURL struct {
	Read  string `json:"read_url"`
	Write string `json:"write_url"`
}

func StructToUrlValue(data interface{}) (url.Values, error) {
	return query.Values(data)
}

type HttpClient struct {
	Client *http.Client
}

func (hc HttpClient) GET(url string) ([]byte, error) {
	handleErr := func(err error) ([]byte, error) {
		return nil, err
	}

	load, err := request(hc.Client, url, "GET", nil)
	if err != nil {
		return handleErr(err)
	}

	return load, nil
}

func (hc HttpClient) POST(url string, load []byte) ([]byte, error) {
	handleErr := func(err error) ([]byte, error) {
		return nil, err
	}

	load, err := request(hc.Client, url, "POST", load)
	if err != nil {
		return handleErr(err)
	}

	return load, nil
}

func (hc HttpClient) PUT(url string, load []byte) ([]byte, error) {
	handleErr := func(err error) ([]byte, error) {
		return nil, err
	}

	load, err := request(hc.Client, url, "PUT", load)
	if err != nil {
		return handleErr(err)
	}

	return load, nil
}

func (hc HttpClient) DELETE(url string, load []byte) ([]byte, error) {
	handleErr := func(err error) ([]byte, error) {
		return nil, err
	}

	load, err := request(hc.Client, url, "DELETE", load)
	if err != nil {
		return handleErr(err)
	}

	return load, nil
}

func request(client *http.Client, url, rtype string, load []byte) ([]byte, error) {
	handleErr := func(err error) ([]byte, error) {
		return nil, fmt.Errorf("http call : %w", err)
	}

	reqBody := bytes.NewBuffer(load)

	req, err := http.NewRequest(rtype, url, reqBody)
	if err != nil {
		return handleErr(fmt.Errorf("prepare request (%w)", err))
	}

	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return handleErr(fmt.Errorf("do request (%w)", err))
	}

	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return handleErr(fmt.Errorf("read response (%w)", err))
	}

	return response, nil
}

func RandomNumber(min, max int) int {
	randomNumberPin := rand.Intn(max-min) + min
	return randomNumberPin
}

//ServiceConfiguration service configuration which to be extract from env vars
type ServiceConfiguration struct {
	App AppConfig
}

//ServiceContext context of service
type ServiceContext struct {
	Config  ServiceConfiguration
	Adaptor RepositoryAdaptor
	Log     *zap.Logger
}

//AppConfig standard configuration for both service and repository
type AppConfig struct {
	Debug      bool
	Timezone   string
	Port       string
	Location   *time.Location `anonymous:"true"`
	LogPath    string
	Name       string
	NetTimeOut time.Duration
}
