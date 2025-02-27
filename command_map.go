package main

import (
	"errors"
	"fmt"

	"github.com/mauricekoreman/go-pokedex/internal/pokecache"
)

func commandMap(cfg *config, cache *pokecache.Cache, parameter string) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL, cache)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationResp.Next
	cfg.previousLocationURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config, cache *pokecache.Cache, parameter string) error {
	if cfg.previousLocationURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationURL, cache)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationResp.Next
	cfg.previousLocationURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
