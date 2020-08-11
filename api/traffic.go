package api

import (
	"clash-cli/model"
	"context"
	"io"

	"github.com/levigross/grequests"
)

func (c *Client) GetTraffics(ctx context.Context) (io.ReadCloser, error) {
	resp, err := grequests.Get(c.BaseURL+model.API_PATH_TRAFFIC, &grequests.RequestOptions{
		Headers: map[string]string{"Authorization": "Bearer " + c.Secret},
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}
	return resp.RawResponse.Body, nil
}
