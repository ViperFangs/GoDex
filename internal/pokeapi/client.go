package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

// Returns a PokeApi Client
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
