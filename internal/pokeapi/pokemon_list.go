package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mauricekoreman/go-pokedex/internal/pokecache"
)

func (c Client) ListPokemon(locationAreaName string, cache *pokecache.Cache) (LocationResponse, error) {
	url := baseURL + "/location-area/" + locationAreaName

	var data []byte
	cachedData, exists := cache.Get(url)
	if !exists {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Error creating request")
			return LocationResponse{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			fmt.Println("Error making request")
			return LocationResponse{}, err
		}

		defer resp.Body.Close()

		d, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response")
			return LocationResponse{}, err
		}

		data = d
	} else {
		data = cachedData
	}
	locationResp := LocationResponse{}
	err := json.Unmarshal(data, &locationResp)
	if err != nil {
		fmt.Println("Error unmarshalling data")
		return LocationResponse{}, err
	}

	return locationResp, nil
}
