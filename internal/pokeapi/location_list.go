package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mauricekoreman/go-pokedex/internal/pokecache"
)

func (c Client) ListLocations(pageURL *string, cache *pokecache.Cache) (Response, error) {
	url := baseURL + "/location-area?offset=0&limit=20"
	if pageURL != nil {
		url = *pageURL
	}

	var data []byte
	cachedData, exists := cache.Get(url)

	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Error creating request")
			return Response{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			fmt.Println("Error making request")
			return Response{}, err
		}

		defer resp.Body.Close()

		d, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response")
			return Response{}, err
		}

		cache.Add(url, d)
		data = d
	} else {
		data = cachedData
	}

	locationsResp := Response{}
	err := json.Unmarshal(data, &locationsResp)
	if err != nil {
		fmt.Println("Error unmarshalling data")
		return Response{}, err
	}

	return locationsResp, nil

}
