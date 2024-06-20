package main

import "fmt"

func callbackHelp(cfg *config, args ...string) error{
	fmt.Println("Welcome to the Pokedex help system.")
	fmt.Println("Here are the available commands:")
	availableCommands := getCommands()
	for _, cmd := range availableCommands{
		fmt.Printf(" - %v \n", cmd.name)
		fmt.Printf("    %v \n", cmd.description)
	}
	fmt.Println("")
	return nil
}