package roundtripper

import (
	"net/http"
	"os"
)

type AuthRoundTripper struct {
	Next http.RoundTripper
}

func (art *AuthRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	token := os.Getenv("TREFLE_TOKEN")

	v := req.URL.Query()
	v.Set("token", token)

	req.URL.RawQuery = v.Encode()
	return art.Next.RoundTrip(req)
}
