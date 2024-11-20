package main

import (
	"errors"
	"fmt"
	"os"
	"math/rand"
)

func commandExit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range getCommands() {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMapf(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("please provide a location name/id")
	}

	location := args[0]
	locationResp, err := cfg.pokeapiClient.ListDetailedLocation(location)
	if err != nil {
		return err
	}

	for _, loc := range locationResp.PokemonEncounters {
		fmt.Println(loc.Pokemon.Name)
	}
	return nil

}

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("please provide a pokemon name/id")
	}

	pokemon := args[0]
	if _, ok := cfg.pokemonCaptured[pokemon]; ok {
		fmt.Printf("You have already caught %s\n", pokemon)
		return nil
	}

	pokemonResp, err := cfg.pokeapiClient.CatchPokemon(pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s, base experience: %d...\n", pokemonResp.Name, pokemonResp.BaseExperience)
	num := rand.Intn(2 * pokemonResp.BaseExperience)

	if num > (pokemonResp.BaseExperience) {
		fmt.Printf("You caught a %s!\n", pokemonResp.Name)
		cfg.pokemonCaptured[pokemonResp.Name] = pokemonResp
	} else {
		fmt.Printf("You failed to catch a %s!\n", pokemonResp.Name)
	}


	return nil
}