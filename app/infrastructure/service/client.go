package service

import (
	"ev-enhance/app/httpapi"
	"net/http"
)

var _ ExternalPartnerAPI = new(Client)

type Client struct {
	getEVRateEndpoint string
	httpClient        httpapi.Caller
}

// NewClient ...
func NewClient(getEVRateEndpoint string) *Client {
	return &Client{
		getEVRateEndpoint: getEVRateEndpoint,
		httpClient:        httpapi.NewCaller(http.DefaultClient),
	}
}

// func (c *Client) SendIntegerToPostmanServer(mockPostmanEndpoint string, result int) error {
//     // Create a JSON payload with the result integer
//     payload := map[string]int{"result": result}
//     jsonPayload, err := json.Marshal(payload)
//     if err != nil {
//         return fmt.Errorf("error marshalling JSON payload: %v", err)
//     }

//     // Create a custom API implementation
//     api := &postmanAPI{endpoint: mockPostmanEndpoint, payload: jsonPayload}

//     // Call the API using the httpClientCaller
//     successful, err := c.httpClientCaller.Call(context.Background(), api)
//     if err != nil {
//         return fmt.Errorf("error calling Postman API: %v", err)
//     }
//     if !successful {
//         return fmt.Errorf("call to Postman API was not successful")
//     }

//     return nil
// }

// // Define a custom API implementation
// type postmanAPI struct {
//     endpoint string
//     payload  []byte
// }

// func (p *postmanAPI) BuildRequest(ctx context.Context) (*http.Request, error) {
//     // Build an HTTP request with the provided endpoint and payload
//     req, err := http.NewRequest("POST", p.endpoint, bytes.NewBuffer(p.payload))
//     if err != nil {
//         return nil, err
//     }
//     req.Header.Set("Content-Type", "application/json")
//     return req, nil
// }

// func (p *postmanAPI) ParseResponse(ctx context.Context, req *http.Request, resp *http.Response) error {
//     // Parse the response if needed
//     // In this example, we are not parsing the response
//     return nil
// }
