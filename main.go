package main

import (
	"github.com/JulianN96/pokedex/internal/pokeapi"
	"time"
)

type config struct {
	//Stateful information goes here
	pokeapiClient pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	caughtPokemon map[string]pokeapi.Pokemon
}


func main(){

	cfg := config {
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),

	}
	startRepl(&cfg)

}