package httpapi

import "net/http"

type DecorateRequestFunc func(req *http.Request) (*http.Request, error)

// DecorateRequest ...
func DecorateRequest(fns ...DecorateRequestFunc) DecorateRequestFunc {
	return func(req *http.Request) (*http.Request, error) {
		for _, fn := range fns {
			var err error
			req, err = fn(req)
			if err != nil {
				return nil, err
			}
		}
		return req, nil
	}
}
