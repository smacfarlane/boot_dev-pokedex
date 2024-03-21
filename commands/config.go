package commands

import "pokedex/internal/pokeapi"

type Config struct {
	Client   pokeapi.Client
	Next     *string
	Previous *string
}
