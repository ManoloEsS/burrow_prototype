package httprequest

import (
	"io"
	"net/http"
	"strings"
)

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
		switch auth.Type {
		case "basic":
			httpRequest.SetBasicAuth(auth.User, auth.Password)
		case "bearer":
			bearer := "Bearer " + auth.BearerToken
			httpRequest.Header.Add("Authorization", bearer)
		}
	}
}
