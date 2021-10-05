package repositoryadaptor

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

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
