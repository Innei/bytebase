// Package oauth provides library for VCS plugins.
package oauth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// TokenRefresher is a function to refresh the OAuth token and assign back to
// the old token upon a successful refresh.
type TokenRefresher func(ctx context.Context, client *http.Client, oldToken *string) error

func requester(ctx context.Context, client *http.Client, method, url string, token *string, body io.Reader) func() (*http.Response, error) {
	return requesterWithHeader(ctx, client, method, url, token, body, nil)
}

func requesterWithHeader(ctx context.Context, client *http.Client, method, url string, token *string, body io.Reader, header map[string]string) func() (*http.Response, error) {
	// The body may be read multiple times but io.Reader is meant to be read once,
	// so we read the body first and build the reader every time.
	var bodyBytes []byte
	if body != nil {
		b, err := io.ReadAll(body)
		if err != nil {
			return func() (*http.Response, error) {
				return nil, errors.Wrap(err, "failed to read from body")
			}
		}
		bodyBytes = b
	}
	return func() (*http.Response, error) {
		req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(bodyBytes))
		if err != nil {
			return nil, errors.Wrapf(err, "construct %s %s", method, url)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *token))
		for k, v := range header {
			req.Header.Set(k, v)
		}

		resp, err := client.Do(req)
		if err != nil {
			return nil, errors.Wrapf(err, "%s %s", method, url)
		}
		return resp, nil
	}
}

// Post makes a HTTP POST request to the given URL using the token. It refreshes
// token and retries the request in the case of the token has expired.
func Post(ctx context.Context, client *http.Client, url string, token *string, body io.Reader, tokenRefresher TokenRefresher) (code int, header http.Header, respBody string, err error) {
	return PostWithHeader(ctx, client, url, token, body, tokenRefresher, nil)
}

// PostWithHeader makes a HTTP POST request to the given URL using the token and
// additional header. It refreshes token and retries the request in the case of
// the token has expired.
func PostWithHeader(ctx context.Context, client *http.Client, url string, token *string, body io.Reader, tokenRefresher TokenRefresher, header map[string]string) (code int, _ http.Header, respBody string, err error) {
	return retry(ctx, client, token, tokenRefresher, requesterWithHeader(ctx, client, http.MethodPost, url, token, body, header))
}

// Get makes a HTTP GET request to the given URL using the token. It refreshes
// token and retries the request in the case of the token has expired.
func Get(ctx context.Context, client *http.Client, url string, token *string, tokenRefresher TokenRefresher) (code int, header http.Header, respBody string, err error) {
	return retry(ctx, client, token, tokenRefresher, requester(ctx, client, http.MethodGet, url, token, nil))
}

// Put makes a HTTP PUT request to the given URL using the token. It refreshes
// token and retries the request in the case of the token has expired.
func Put(ctx context.Context, client *http.Client, url string, token *string, body io.Reader, tokenRefresher TokenRefresher) (code int, header http.Header, respBody string, err error) {
	return retry(ctx, client, token, tokenRefresher, requester(ctx, client, http.MethodPut, url, token, body))
}

// Patch makes a HTTP PATCH request to the given URL using the token. It
// refreshes token and retries the request in the case of the token has expired.
func Patch(ctx context.Context, client *http.Client, url string, token *string, body io.Reader, tokenRefresher TokenRefresher) (code int, header http.Header, respBody string, err error) {
	return retry(ctx, client, token, tokenRefresher, requester(ctx, client, http.MethodPatch, url, token, body))
}

// Delete makes a HTTP DELETE request to the given URL using the token. It refreshes
// token and retries the request in the case of the token has expired.
func Delete(ctx context.Context, client *http.Client, url string, token *string, tokenRefresher TokenRefresher) (code int, header http.Header, respBody string, err error) {
	return retry(ctx, client, token, tokenRefresher, requester(ctx, client, http.MethodDelete, url, token, nil))
}

const maxRetries = 3

func retry(ctx context.Context, client *http.Client, token *string, tokenRefresher TokenRefresher, f func() (*http.Response, error)) (code int, header http.Header, respBody string, err error) {
	var resp *http.Response
	var body []byte
	for retries := 0; retries < maxRetries; retries++ {
		select {
		case <-ctx.Done():
			return 0, nil, "", ctx.Err()
		default:
		}

		resp, err = f()
		if err != nil {
			return 0, nil, "", err
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return 0, nil, "", errors.Wrapf(err, "read response body with status code %d", resp.StatusCode)
		}

		if err = getOAuthErrorDetails(resp.StatusCode, body); err != nil {
			if _, ok := err.(*oauthError); ok {
				// Refresh the token
				if err := tokenRefresher(ctx, client, token); err != nil {
					return 0, nil, "", err
				}
				continue
			}
			return 0, nil, "", errors.Errorf("got unexpected OAuth error %T", err)
		}
		return resp.StatusCode, resp.Header, string(body), nil
	}
	return 0, nil, "", errors.Errorf("retries exceeded for OAuth refresher with status code %d and body %q", resp.StatusCode, string(body))
}

type oauthError struct {
	Err              string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (e oauthError) Error() string {
	return fmt.Sprintf("OAuth response error %q description %q", e.Err, e.ErrorDescription)
}

// getOAuthErrorDetails only returns error if it's an OAuth error. For other
// errors like 404 we don't return error. We do this because this method is only
// intended to be used by oauth to refresh access token on expiration.
//
// When it's error like 404, GitLab API doesn't return it as error so we keep
// the similar behavior and let caller check the response status code.
func getOAuthErrorDetails(code int, body []byte) error {
	if 200 <= code && code < 300 {
		return nil
	}

	// Special case for Bitbucket OAuth error.
	if bytes.Contains(body, []byte("OAuth2 access token expired.")) {
		return &oauthError{}
	}

	var oe oauthError
	if err := json.Unmarshal(body, &oe); err != nil {
		// If we failed to unmarshal body with oauth error, it's not oauthError and we should return nil.
		// nolint:nilerr
		return nil
	}
	// https://www.oauth.com/oauth2-servers/access-tokens/access-token-response/
	// {"error":"invalid_token","error_description":"Token is expired. You can either do re-authorization or token refresh."}
	// {"error":"invalid_grant","error_description":"The provided authorization grant is invalid, expired, revoked, does not match the redirection URI used in the authorization request, or was issued to another client."}
	if (oe.Err == "invalid_token" || oe.Err == "invalid_grant") && strings.Contains(oe.ErrorDescription, "expired") {
		return &oe
	}
	return nil
}
