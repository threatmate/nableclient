package ncentralclient

import (
	"context"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/threatmate/restapiclient"
)

// ClientOption is a function that configures the client.
type ClientOption func(*Config)

// Config is the configuration for the client.
type Config struct {
	BaseURL    string
	HTTPClient *http.Client
}

// Client is the client for the n-able REST API.
type Client struct {
	client      *restapiclient.Client
	lock        sync.Mutex
	accessToken string
}

// New creates a new client for the n-able REST API.
//
// The default HTTP client under the hood will have a 1-minute timeout.
func New(baseURL string, opts ...ClientOption) *Client {
	config := Config{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 1 * time.Minute,
		},
	}

	for _, opt := range opts {
		opt(&config)
	}

	if !strings.Contains(config.BaseURL, "://") {
		config.BaseURL = "https://" + config.BaseURL
	}

	client := restapiclient.New(config.BaseURL)
	if config.HTTPClient != nil {
		httpClient := client.HTTPClient()
		*httpClient = *config.HTTPClient
	}

	return &Client{
		client: client,
	}
}

// GetClient returns the underlying REST API client.
func (c *Client) HTTPClient() *http.Client {
	if c.client == nil {
		return nil
	}
	return c.client.HTTPClient()
}

// Do performs a request to the n-able REST API.
func (c *Client) Do(ctx context.Context, method string, path string, input any, output any, options ...restapiclient.Option) error {
	var accessToken string
	c.lock.Lock()
	accessToken = c.accessToken
	c.lock.Unlock()

	var newOptions []restapiclient.Option
	if accessToken != "" {
		newOptions = append(newOptions, restapiclient.OptionHeader("Authorization", "Bearer "+accessToken))
	}
	newOptions = append(newOptions, options...)

	return c.client.Do(ctx, method, path, input, output, newOptions...)
}
