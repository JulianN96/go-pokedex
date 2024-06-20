package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//Generate a function built off of the client struct to get the LocationAreas
func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	//List the custom endpoint
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	// Check the cache. If cache exists, return the cached data rather than make an API call
	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit!")
		pokemonName := Pokemon{}
		err := json.Unmarshal(data, &pokemonName)
		if err != nil {
			return Pokemon{}, err
		}
	
		//Return the new struct and nil error
		return pokemonName, nil
	}
	fmt.Println("cache miss")

	//Generate a reqest and error and make the request using http.NewRequest("METHOD", URL, ERROR)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}
	//Assign a response variable using c.httpClient.Do(req)
	resp, err := c.httpClient.Do(req)
	if err != nil{
		return Pokemon{}, err
	}
	//Check for problem error codes
	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	//Create a data variable by reading the response stream
	dat, err := io.ReadAll(resp.Body)

	if err != nil {
		return Pokemon{}, err
	}

	//Generate a new version of the locations struct
	pokemonNam := Pokemon{}

	//Pass the json data into the struct. If it has an error it will generate
	err = json.Unmarshal(dat, &pokemonNam)
	if err != nil {
		return Pokemon{}, err
	}

	//Add new items to the cache
	c.cache.Add(fullURL, dat)

	//Return the new struct and nil error
	return pokemonNam, nil
}