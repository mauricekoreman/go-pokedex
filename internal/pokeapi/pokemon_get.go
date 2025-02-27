package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) GetPokemon(pokemonName string) (PokemonResponse, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if cachedData, exists := c.cache.Get(url); exists {
		pokemonResp := PokemonResponse{}
		err := json.Unmarshal(cachedData, &pokemonResp)
		if err != nil {
			return PokemonResponse{}, err
		}

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request")
		return PokemonResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println("Error making request")
		return PokemonResponse{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response")
		return PokemonResponse{}, err
	}

	pokemonResp := PokemonResponse{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		fmt.Println("Error unmarshalling data")
		return PokemonResponse{}, err
	}

	return pokemonResp, nil
}
