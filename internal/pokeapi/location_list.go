package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (locationResponse, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	data, ok := c.cache.Get(url)
	if ok {
		return unmarshalLocationResponse(data)
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

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return locationResponse{}, err
	}

	c.cache.Add(url, data)

	return unmarshalLocationResponse(data)

}

func unmarshalLocationResponse(data []byte) (locationResponse, error) {
	locationsResponse := locationResponse{}
	err := json.Unmarshal(data, &locationsResponse)
	if err != nil {
		return locationResponse{}, err
	}

	return locationsResponse, nil
}
