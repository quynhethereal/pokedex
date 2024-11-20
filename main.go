package main

import (
	"time"

	"github.com/soapycattt/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	pokemonCaptured := make(map[string]pokeapi.Pokemon)

	cfg := &config{
		pokeapiClient:    pokeClient,
		pokemonCaptured: pokemonCaptured,
	}

	startRepl(cfg)
}