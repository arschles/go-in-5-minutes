package dsclient

import (
	"fmt"

	multierr "github.com/hashicorp/go-multierror"
	"github.com/parnurzeal/gorequest"
)

type Client struct {
	apiKey string
	cl     *gorequest.SuperAgent
}

func New(apiKey string) (*Client, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("No API key passed")
	}
	return &Client{apiKey: apiKey, cl: gorequest.New()}, nil
}

func (c *Client) Forecast(lat, long float64) (*Response, error) {
	forecastURL := fmt.Sprintf(
		"https://api.darksky.net/forecast/%s/%f,%f",
		c.apiKey,
		lat,
		long,
	)
	resp := &Response{}
	httpRes, _, errs := c.cl.Get(forecastURL).EndStruct(resp)
	if len(errs) > 0 {
		return nil, &multierr.Error{Errors: errs}
	}
	if httpRes.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Status Code %d returned!", httpRes.StatusCode)
	}
	return resp, nil
}
