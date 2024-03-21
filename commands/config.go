package commands

import (
	"pokedex/internal/pokeapi"
	"pokedex/internal/pokedex"
)

type Config struct {
	Client   pokeapi.Client
	Pokedex  pokedex.Pokedex
	Next     *string
	Previous *string
}
