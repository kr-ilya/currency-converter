package provider

import (
	"context"
	"currency-telegram-webapp-backend/internal/client"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type Provider interface {
	Latest(ctx context.Context, attr ...url.Values) (*FixerResponse, error)
}

type FixedProvider struct {
	api         *client.Caller
	baseURL     *url.URL
	accessToken string
}

type Currency string

type Rates map[Currency]float64

type Date struct {
	time.Time
}

type FixerResponse struct {
	Success   bool     `json:"success"`
	Timestamp int64    `json:"timestamp,omitempty"`
	Base      Currency `json:"base,omitempty"`
	Rates     Rates    `json:"rates,omitempty"`
}

func NewFixedProvider(api *client.Caller, FixerAccessToken string) *FixedProvider {
	return &FixedProvider{
		api: api,
		baseURL: &url.URL{
			Scheme: "http",
			Host:   "data.fixer.io",
			Path:   "/api",
		},
		accessToken: FixerAccessToken,
	}
}

func Base(c Currency) url.Values {
	v := url.Values{}

	if s := string(c); s != "" {
		v.Set("base", s)
	}

	return v
}

func (p *FixedProvider) query(attr []url.Values) url.Values {
	v := url.Values{}

	for _, a := range attr {
		if base := a.Get("base"); base != "" {
			v.Set("base", base)
		}
	}

	return v
}

func (p *FixedProvider) doGet(ctx context.Context, path string, query url.Values) ([]byte, error) {
	link := p.baseURL.Path + path

	query.Set("access_key", p.accessToken)

	link += "?" + query.Encode()

	rUrl, err := url.Parse(link)
	if err != nil {
		return nil, fmt.Errorf("url parse: %w", err)
	}

	return p.api.Get(p.baseURL.ResolveReference(rUrl).String())
}

func (p *FixedProvider) Latest(ctx context.Context, attr ...url.Values) (*FixerResponse, error) {
	resp, err := p.doGet(ctx, "/latest", p.query(attr))
	if err != nil {
		return nil, err
	}

	result := &FixerResponse{}

	err = json.Unmarshal(resp, result)
	if err != nil {
		return nil, fmt.Errorf("decode json: %w", err)
	}

	return result, nil
}
