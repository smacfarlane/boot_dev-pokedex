package pokeapi

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"pokedex/internal/pokecache"
	"time"
)

const baseUrl string = "https://pokeapi.co/api/v2/"

type Client struct {
	cache pokecache.Cache
}

func NewClient() Client {
	return Client{cache: pokecache.NewCache(time.Hour)}
}

func (c *Client) Get(url *string) ([]byte, error) {
	// Return from cache if it exists
	if b, ok := c.cache.Get(*url); ok {
		return b, nil
	}

	res, err := http.Get(*url)
	if err != nil {
		return []byte{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		err := fmt.Sprintf("server returned %d", res.StatusCode)
		return []byte{}, errors.New(err)
	}
	if err != nil {
		return []byte{}, err
	}

	// Add response to cache
	c.cache.Add(*url, body)

	return body, nil
}
