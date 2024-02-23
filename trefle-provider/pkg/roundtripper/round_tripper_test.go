package roundtripper

import (
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockRoundTripper struct {
	req *http.Request
}

func (mrt *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	mrt.req = req
	return &http.Response{StatusCode: 200}, nil
}

func TestRoundTrip(t *testing.T) {
	assert.NoError(t, os.Setenv("TREFLE_TOKEN", "abc"))

	var (
		req = http.Request{URL: &url.URL{}}
		art = AuthRoundTripper{Next: &mockRoundTripper{}}
	)

	res, err := art.RoundTrip(&req)

	assert.Nil(t, err)
	assert.Equal(t, &http.Response{StatusCode: 200}, res)
	assert.Equal(
		t,
		req.URL.Query(),
		url.Values{"token": []string{"abc"}},
	)
}
