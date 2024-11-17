package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func([]string) error
}

func commandHelp(args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("help - Displays a help message")
	fmt.Println("exit - Exit the Pokedex")
	return nil
}

func commandExit(args []string) error {
	fmt.Println("Exiting the Pokedex...")
	os.Exit(0)
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func main() {
	for {
		fmt.Print("Pokedex > ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		commandMap := getCommands()
		args := strings.Split(input, " ")
		command, ok := commandMap[args[0]]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		err := command.callback(args[1:])
		if err != nil {
			fmt.Println(err)
		}

	}
}
