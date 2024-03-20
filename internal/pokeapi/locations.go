package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// TODO Rename?
type BasicResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(pagedUrl *string) (BasicResponse, error) {
	url := baseUrl + "/location-area"
	if pagedUrl != nil {
		url = *pagedUrl
	}

	res, err := http.Get(url)
	if err != nil {
		return BasicResponse{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return BasicResponse{}, errors.New("error from server") // TODO: Propogate the code?
	}
	if err != nil {
		return BasicResponse{}, err
	}

	loc := BasicResponse{}

	err = json.Unmarshal(body, &loc)
	if err != nil {
		return BasicResponse{}, err
	}
	return loc, nil
}
