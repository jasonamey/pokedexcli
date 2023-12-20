package main

import (
	"fmt"
)

func callbackHelp() error {
	fmt.Println("Welcome to the Pokedex help menu")
	fmt.Println("Here are your commands")
	commands := getCommands()

	for _, command := range commands {
		fmt.Printf(" - %s: %s\n", command.name, command.description)
	}
	// fmt.Println(" - help")
	// fmt.Println(" - exit")
	fmt.Println("")
	return nil
}
