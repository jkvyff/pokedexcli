package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("-------Simple Pokedex-------")
	fmt.Print(" > ")

	for scanner.Scan() {

		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}
		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}

}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Attempt to catch a pokemon and add it to your pokedex",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "View information about a caught pokemon",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View all the pokemon in your pokedex",
			callback:    callbackPokedex,
		},
		"map": {
			name:        "map",
			description: "Provides the next page of areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Provides the previous page of areas",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore {area_name}",
			description: "Lists the pokemon in a location area",
			callback:    callbackExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
