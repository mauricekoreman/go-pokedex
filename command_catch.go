package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemonResp, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	baseExp := pokemonResp.BaseExperience
	randomNum := rand.Intn(baseExp)

	// 1 in 3 chance of catching the pokemon
	if randomNum%3 == 0 {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.caughtPokemon[pokemonName] = pokemonResp
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}
