package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//Generate a function built off of the client struct to get the LocationAreas
func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	//List the custom endpoint
	endpoint := "/location-area/"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	// Check the cache. If cache exists, return the cached data rather than make an API call
	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit!")
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(data, &locationAreasResp)
		if err != nil {
			return LocationAreasResp{}, err
		}
	
		//Return the new struct and nil error
		return locationAreasResp, nil
	}
	fmt.Println("cache miss")

	//Generate a reqest and error and make the request using http.NewRequest("METHOD", URL, ERROR)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}
	//Assign a response variable using c.httpClient.Do(req)
	resp, err := c.httpClient.Do(req)
	if err != nil{
		return LocationAreasResp{}, err
	}
	//Check for problem error codes
	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	//Create a data variable by reading the response stream
	dat, err := io.ReadAll(resp.Body)

	if err != nil {
		return LocationAreasResp{}, err
	}

	//Generate a new version of the locations struct
	locationAreasResp := LocationAreasResp{}

	//Pass the json data into the struct. If it has an error it will generate
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	//Add new items to the cache
	c.cache.Add(fullURL, dat)

	//Return the new struct and nil error
	return locationAreasResp, nil
}


//Generate a function built off of the client struct to get the LocationAreas
func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	//List the custom endpoint
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	// Check the cache. If cache exists, return the cached data rather than make an API call
	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit!")
		locationArea := LocationArea{}
		err := json.Unmarshal(data, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
	
		//Return the new struct and nil error
		return locationArea, nil
	}
	fmt.Println("cache miss")

	//Generate a reqest and error and make the request using http.NewRequest("METHOD", URL, ERROR)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}
	//Assign a response variable using c.httpClient.Do(req)
	resp, err := c.httpClient.Do(req)
	if err != nil{
		return LocationArea{}, err
	}
	//Check for problem error codes
	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	//Create a data variable by reading the response stream
	dat, err := io.ReadAll(resp.Body)

	if err != nil {
		return LocationArea{}, err
	}

	//Generate a new version of the locations struct
	locationArea := LocationArea{}

	//Pass the json data into the struct. If it has an error it will generate
	err = json.Unmarshal(dat, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	//Add new items to the cache
	c.cache.Add(fullURL, dat)

	//Return the new struct and nil error
	return locationArea, nil
}