package main

import (
	"fmt"
	"errors"
)

func callbackMap(cfg *config, args ...string) error{
	
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results{
		fmt.Printf(" - %v\n", area.Name)
	}
	//Every time we get a new page, the next and previous pointers are updated
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}

func callbackMapb(cfg *config, args ...string) error{
	if cfg.prevLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results{
		fmt.Printf(" - %v\n", area.Name)
	}
	//Every time we get a new page, the next and previous pointers are updated
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}