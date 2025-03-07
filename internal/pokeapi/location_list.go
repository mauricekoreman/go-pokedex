package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) ListLocations(pageURL *string) (LocationsResponse, error) {
	url := baseURL + "/location-area?offset=0&limit=20"
	if pageURL != nil {
		url = *pageURL
	}

	if cachedData, exists := c.cache.Get(url); exists {
		locationResp := LocationsResponse{}
		err := json.Unmarshal(cachedData, &locationResp)
		if err != nil {
			return LocationsResponse{}, err
		}

		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request")
		return LocationsResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println("Error making request")
		return LocationsResponse{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response")
		return LocationsResponse{}, err
	}

	locationsResp := LocationsResponse{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		fmt.Println("Error unmarshalling data")
		return LocationsResponse{}, err
	}

	c.cache.Add(url, data)
	return locationsResp, nil

}
