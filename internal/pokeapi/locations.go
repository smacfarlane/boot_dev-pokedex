package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
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

func GetLocationAreas(pagedUrl *string) (BasicResponse, error) {
	url := baseUrl + "location-area/"
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
		err := fmt.Sprintf("server returned %d", res.StatusCode)
		return BasicResponse{}, errors.New(err)
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

func GetLocationArea(name string) (LocationArea, error) {
	url := baseUrl + "location-area/" + name
	res, err := http.Get(url)
	if err != nil {
		return LocationArea{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		err := fmt.Sprintf("server returned %d", res.StatusCode)
		return LocationArea{}, errors.New(err)
	}
	if err != nil {
		return LocationArea{}, err
	}

	loc := LocationArea{}

	err = json.Unmarshal(body, &loc)
	if err != nil {
		return LocationArea{}, err
	}
	return loc, nil
}
