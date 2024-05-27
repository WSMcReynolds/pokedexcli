package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (locationResponse, error) {
	url := baseURL + "/location"

	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationResponse{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationResponse{}, err
	}

	locationsResponse := locationResponse{}
	err = json.Unmarshal(data, &locationsResponse)
	if err != nil {
		return locationResponse{}, err
	}

	return locationsResponse, nil

}
