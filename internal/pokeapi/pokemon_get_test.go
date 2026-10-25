package pokeapi

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func newTestClient(rt roundTripFunc) Client {
	c := NewClient(time.Second, time.Minute)
	c.httpClient.Transport = rt
	return c
}

func TestGetPokemonErrors(t *testing.T) {
	cases := map[string]roundTripFunc{
		"network error": func(*http.Request) (*http.Response, error) {
			return nil, errors.New("connection refused")
		},
		"pokemon not found": func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusNotFound,
				Body:       io.NopCloser(strings.NewReader("Not Found")),
				Request:    req,
			}, nil
		},
	}

	for name, rt := range cases {
		t.Run(name, func(t *testing.T) {
			c := newTestClient(rt)
			_, err := c.GetPokemon("missingno")
			if err == nil {
				t.Fatalf("expected an error, got nil")
			}
		})
	}
}
