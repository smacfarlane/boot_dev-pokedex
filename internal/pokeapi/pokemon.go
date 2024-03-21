package pokeapi

import (
	"encoding/json"
)

type Pokemon struct {
	Id             int        `json:"id"`
	Name           string     `json:"name"`
	BaseExperience int        `json:"base_experience"`
	Height         int        `json:"height"`
	IsDefault      bool       `json:"is_default"`
	Order          int        `json:"order"`
	Weight         int        `json:"weight"`
	Abilities      []struct{} `json:"abilities"`
	Forms          []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"forms"`
	GameIndicies           []struct{} `json:"game_indicies"`
	HeldItems              []struct{} `json:"held_items"`
	LocationAreaEncounters string     `json:"location_area_encounters"`
	Moves                  []struct{} `json:"moves"`
	Species                struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"species"`
	Sprites struct{} `json:"sprites"`
	Cries   struct{} `json:"cries"`
	Stats   []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		}
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		}
	} `json:"types"`
	PastTypes []struct{} `json:"past_types"`
}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseUrl + "pokemon/" + name

	response, err := c.Get(&url)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(response, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
