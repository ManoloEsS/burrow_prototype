package httprequest

import (
	"context"
	"fmt"
	"net/http"
)

func httpGetRequest(ctx context.Context, client *http.Client, reqParams *UserRequest) (*http.Response, error) {

	req, err := http.NewRequestWithContext(ctx, reqParams.Method, reqParams.URL, nil)
	if err != nil {
		return nil, fmt.Errorf("Couldn't create request: %w", err)
	}

	addRequestParameters(req, reqParams.Params)

	addRequestHeaders(req, reqParams.Headers)

	addRequestAuth(req, &reqParams.Auth)

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Couldn't get response: %w", err)
	}

	return res, nil
}
