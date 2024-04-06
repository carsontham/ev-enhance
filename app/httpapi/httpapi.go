package httpapi

import (
	"context"
	"fmt"
	"net/http"
)

type (
	API interface {
		BuildRequest(ctx context.Context) (*http.Request, error)
		ParseResponse(ctx context.Context, req *http.Request, resp *http.Response) error
	}

	// Caller ...
	Caller interface {
		Call(ctx context.Context, api API) (successful bool, err error)
	}
)

// NewCaller ...
func NewCaller(c *http.Client) Caller {
	return &caller{c}
}

// Caller ...
type caller struct {
	httpClient *http.Client
}

// Call ...
func (c caller) Call(ctx context.Context, api API) (successful bool, err error) {
	req, err := api.BuildRequest(ctx)
	if err != nil {
		return false, err
	}

	req = req.WithContext(ctx)

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return false, err
	}
	defer func() {
		if resp.Body != nil {
			err := resp.Body.Close()
			fmt.Println(err)
		}
	}()

	if err := api.ParseResponse(ctx, req, resp); err != nil {
		return false, err
	}

	return true, nil
}
