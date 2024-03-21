package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	
	req, err := http.NewRequest("Get", fullURL, nil)
	if err != nil{
		return LocationAreasResp{}, err
	}
	
	resp, err := c.httpClient.Do(req)
	if err != nil{
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()
	fmt.Println(resp)
	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreasResp, nil
}