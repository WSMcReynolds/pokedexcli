package pokeapi

import (
	"encoding/json"
	"io"
)

func (c *Client) GetAreaInfo(area string) (areaResponse, error) {

	areaURL := baseURL + "/location-area/" + area

	val, ok := c.cache.Get(areaURL)
	res := areaResponse{}

	if ok {
		err := json.Unmarshal(val, &res)

		if err != nil {
			return areaResponse{}, err
		}

		return res, nil
	}

	resp, err := c.httpClient.Get(areaURL)
	if err != nil {
		return areaResponse{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return areaResponse{}, err
	}

	c.cache.Add(areaURL, data)

	err = json.Unmarshal(data, &res)

	if err != nil {
		return areaResponse{}, err
	}

	return res, nil
}
