package httprequest

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
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

func getBodyContents(reqBody Body) io.Reader {
	switch reqBody.Type {
	case "text":
		return strings.NewReader(reqBody.Content)
	default:
		return nil
	}
}

func addRequestParameters(httpRequest *http.Request, params map[string]string) {
	if len(params) > 0 {
		q := httpRequest.URL.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		httpRequest.URL.RawQuery = q.Encode()
	}
}

func addRequestHeaders(httpRequest *http.Request, headers map[string]string) {
	if len(headers) > 0 {
		for k, v := range headers {
			httpRequest.Header.Add(k, v)
		}
	}
}

func addRequestAuth(httpRequest *http.Request, auth *Auth) {
	if auth != nil {

	}

}
