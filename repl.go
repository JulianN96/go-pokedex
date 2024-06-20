package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)


func startRepl(cfg *config){
	fmt.Println("Pokedex online.")
	//Generate a scanner
	scanner := bufio.NewScanner(os.Stdin)
	

	for {
		fmt.Print("Please enter your command:")
		//Scan the input from the os
		scanner.Scan()
		text := scanner.Text()
		//Convert command to lowercase command
		cleaned := cleanInput(text)
		//If no command present, restart loop
		if len(cleaned) == 0{
			continue
		}
		//Assign the cleaned command a variable
		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:] 
		}

		//Get all available commands
		availableCommands := getCommands()
		
		//If the command exists, call it.
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
		}
		err := command.callback(cfg, args...)
		if err != nil{
			fmt.Println(err)
		}
	}
}

type cliCommand struct{
	name string
	description string
	callback func(*config, ...string) error
}

func getCommands() map[string]cliCommand{
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays the Help Menu with all available commands",
			callback: callbackHelp,
		},
		"exit": {
			name: "exit",
			description: "Exits the program",
			callback: callbackExit,
		},
		"map": {
			name: "map",
			description: "Lists next page of areas",
			callback: callbackMap,
		},
		"mapb": {
			name: "mapb",
			description: "Lists previous page of areas",
			callback: callbackMapb,
		},
		"explore": {
			name: "explore {locationArea}",
			description: "Lists all pokemon in that area",
			callback: callbackExplore,
		},
		"catch": {
			name: "catch {pokemonName}",
			description: "Attempt to catch a pokemon",
			callback: callbackCatch,
		},
		"inspect": {
			name: "inspect {pokemonName}",
			description: "View information about caught pokemon",
			callback: callbackInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "List all caught pokemon",
			callback: callbackPokedex,
		},
		
	}
}

func cleanInput (str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}