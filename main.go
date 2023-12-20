package main

import "fmt"

func commandHelp() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage: \n\n")
	fmt.Println("help: displays a help message")
	fmt.Println("exit: exits the Pokedex")

}

func commandExit() {

}

type CommandFunction func()

type cliCommand struct {
	name        string
	description string
	callback    func()
}

func helpCli() {

}

func exitCli() {

}

var commands = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "doing the help",
		callback:    helpCli,
	},
	"exit": {
		name:        "exit",
		description: "goodbye",
		callback:    exitCli,
	},
}

func main() {
	program_continue := true
	for program_continue {
		var response string
		fmt.Print("Pokedex >")
		fmt.Scan(&response)

	}
}
