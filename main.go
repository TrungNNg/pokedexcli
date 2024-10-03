package main

import (
	"fmt"
	"os"
)

var p = fmt.Println

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	message := `

Wellcome to the Pokedex!
Usage:

help: Display a help message
exit: Exit the Pokedex
    
`
	p(message)
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func main() {

	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	input := ""
	for {
		fmt.Print("Pokedex > ")
		fmt.Scan(&input)
		//fmt.Println(input)
		if cmd, ok := commands[input]; ok {
			cmd.callback()
		} else {
			p("Invalid command, use \"help\" for list of commands.")
		}
	}

}
