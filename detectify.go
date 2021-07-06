package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	BaseEndpoint     = "api.detectify.com/rest/v2"
	APIKeyHeaderName = "X-Detectify-Key"
)

// Provider is a struct containing an http Client so that we can communicate with Detectify's public API
type Provider struct {
	client http.Client
}

// NewProvider initializes and returns a new Provider instance with the specified http Client timeout
func NewProvider(timeout time.Duration) *Provider {
	return &Provider{
		client: http.Client{
			Timeout: timeout,
		},
	}
}

// StartScan starts a scan on the specified scan profile token
func (p *Provider) StartScan() (string, error) {
	path := fmt.Sprintf("scans/%s/", ScanProfileToken)

	response, err := p.sendRequest(path)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (p *Provider) sendRequest(path string) (string, error) {
	url := fmt.Sprintf("https://%s/%s", BaseEndpoint, path)
	req, err := http.NewRequest(
		http.MethodPost,
		url,
		nil,
	)
	if err != nil {
		log.Println(err)
		return "", err
	}
	req.Header.Set(APIKeyHeaderName, ApiKey)

	response, err := p.client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(body), nil
}
