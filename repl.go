package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		command, ok := getCommands()[cleaned[0]]

		if !ok {
			fmt.Println("invalid command!")
			continue
		}

		command.callback()

		// switch command.name {
		// case "help":
		// 	command.callback()
		// 	// fmt.Println("Welcome to the Pokedex help menu")
		// 	// fmt.Println("Here are your commands")
		// 	// fmt.Println(" - help")
		// 	// fmt.Println(" - exit")
		// 	// fmt.Println("")

		// case "exit":
		// 	command.callback()
		// default:
		// 	fmt.Println("Invalid Command")
		// }
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Turns off the Pokedex",
			callback:    exitHelp,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)

	return words
}
