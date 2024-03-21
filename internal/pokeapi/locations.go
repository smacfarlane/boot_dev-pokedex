package pokeapi

import (
	"encoding/json"
	"fmt"
)

type BasicResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

type LocationArea struct {
	EncounterMethodRates []struct{} `json:"encounter_method_rates"`
	Location             struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"location"`
	Names      []struct{} `json:"names"`
	Encounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct{} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetLocationAreas(pagedUrl *string) (BasicResponse, error) {
	url := baseUrl + "location-area/"
	if pagedUrl != nil {
		url = *pagedUrl
	}

	response, err := c.Get(&url)
	if err != nil {
		return BasicResponse{}, err
	}

	loc := BasicResponse{}

	err = json.Unmarshal(response, &loc)
	if err != nil {
		return BasicResponse{}, err
	}
	fmt.Printf("%v\n", loc)
	return loc, nil
}

func (c *Client) GetLocationArea(name string) (LocationArea, error) {
	url := baseUrl + "location-area/" + name

	response, err := c.Get(&url)
	if err != nil {
		return LocationArea{}, err
	}
	loc := LocationArea{}

	err = json.Unmarshal(response, &loc)
	if err != nil {
		return LocationArea{}, err
	}
	return loc, nil
}
