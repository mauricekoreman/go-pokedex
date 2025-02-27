package main

import (
	"fmt"

	"github.com/mauricekoreman/go-pokedex/internal/pokecache"
)

func commandExplore(cfg *config, cache *pokecache.Cache, locationAreaName string) error {
	fmt.Printf("Exploring %s...\n", locationAreaName)
	listPokemonResp, err := cfg.pokeapiClient.ListPokemon(locationAreaName, cache)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range listPokemonResp.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
