// ListLocations - needs to pass string
// returns Locations and error
package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"

	if pageUrl != nil {
		url = *pageUrl
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, nil
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return RespShallowLocations{}, nil
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)

	if err != nil {
		return RespShallowLocations{}, nil
	}

	locationsResp := RespShallowLocations{}

	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, nil
	}

	return locationsResp, nil
}
