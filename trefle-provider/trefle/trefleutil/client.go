package trefleutil

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	urlFmt       = "https://trefle.io/api/v1/%s"
	urlSearchFmt = "https://trefle.io/api/v1/%s/search"
)

type Config struct {
	Client *http.Client
	Entity string
}

type Client struct {
	Config
}

func (c *Client) Do(ctx context.Context, method, path string, body io.Reader, opts ...Option) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, path, body)

	if err != nil {
		return nil, err
	}

	vs := make(url.Values)

	for _, opt := range opts {
		opt.Set(vs)
	}

	req.URL.RawQuery = vs.Encode()
	req.Header.Set("Content-Type", "application/json")

	return c.Client.Do(req)
}

func (c *Client) List(ctx context.Context, opts ...Option) (json.RawMessage, error) {
	resp, err := c.Do(
		ctx,
		http.MethodGet,
		fmt.Sprintf(urlFmt, c.Entity),
		http.NoBody,
		opts...,
	)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, nil // TODO
	}

	var p json.RawMessage

	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return nil, err
	}

	return p, nil
}

func (c *Client) Search(ctx context.Context, opts ...Option) (json.RawMessage, error) {
	resp, err := c.Do(
		ctx,
		http.MethodGet,
		fmt.Sprintf(urlSearchFmt, c.Entity),
		http.NoBody,
		opts...,
	)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, nil // TODO
	}

	var p json.RawMessage

	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return nil, err
	}

	return p, nil
}
