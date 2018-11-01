// Package http provides a wrapper for http functionality.
package http

import (
	"github.com/pkg/errors"
	"net/http"
)

// Client is really just a wrapper for net/http.Client.
type Client struct {
	httpClient *http.Client
}

// Getter is a simple HTTP GET interface.
type Getter interface {
	Get(url string) (*http.Response, error)
}

// NewClient creates a new Client.
func NewClient(httpClient *http.Client) *Client {
	return &Client{httpClient: httpClient}
}

// Get issues a GET to the specified URL.
func (c Client) Get(url string) (*http.Response, error) {
	response, err := c.httpClient.Get(url)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read from URL %s", url)
	}
	return response, nil
}
