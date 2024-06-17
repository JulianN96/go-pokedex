package main

import (
	"fmt"
)



func main(){
	fmt.Println("Pokedex online.")

	type cliCommand struct {
		name string
		description string
		callback func() error
	}

	func commandHelp(){
		fmt.Println("HelpCommand")
	}

	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		
	}

}