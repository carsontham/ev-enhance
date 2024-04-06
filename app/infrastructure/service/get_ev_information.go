package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"ev-enhance/app/httpapi"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetEVInformation(ctx context.Context, id int) (*EVResponse, error) {

	evRequestForm := EVRequestForm{
		operator: id,
	}


	api := &getEVInformationAPI{
		endpoint: c.getEVRateEndpoint,
		decorReq: httpapi.DecorateRequest(),
		req:      evRequestForm,
	}

	successful, err := c.httpClient.Call(ctx, api)
	if successful {
		return api.resp, nil
	} else {
		fmt.Println("error in calling")
		return nil, err
	}
}

var _ httpapi.API = new(getEVInformationAPI)

type getEVInformationAPI struct {
	endpoint string
	decorReq httpapi.DecorateRequestFunc
	req      EVRequestForm
	resp     *EVResponse
}

func (api *getEVInformationAPI) BuildRequest(ctx context.Context) (*http.Request, error) {
	body, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, api.endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	return api.decorReq(req)
}

func (api *getEVInformationAPI) ParseResponse(_ context.Context, _ *http.Request, resp *http.Response) error {
	var respData EVResponse
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%w, statusCode: %d, body: %s", errors.New("status not ok"), resp.StatusCode, resp.Body)
	}
	body, err := io.ReadAll(resp.Body)
	h := string(body)
	fmt.Println(h)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &respData); err != nil {
		return err
	}

	api.resp = &respData
	return nil
}
