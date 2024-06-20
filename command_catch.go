package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error{

	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]
	
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	randNum := rand.Intn(pokemon.BaseExperience)
	const threshold = 50
	
	if randNum > threshold {
		return fmt.Errorf("Failed to catch %v", pokemonName)
	}

	cfg.caughtPokemon[pokemonName] = pokemon

	fmt.Printf("%v was caught!\n", pokemonName)

	return nil
}