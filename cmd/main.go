package main

import (
	"bufio"
	"fmt"
	"os"

	pokedex "github.com/anantashahane/pokedex/pokedex"
)

type cliCommand struct {
	name        string
	description string
	callback    func(configuration *config) error
}

type config struct {
	previous string
	variable string
	next     string
}

func commandExit(configuration *config) error {
	fmt.Println("\tClosing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("\tError quiting")
}

func makeHelpCommand(commands map[string]cliCommand) func(configuration *config) error {
	return func(configuration *config) error {
		fmt.Println("\tWelcome to the Pokedex!")
		fmt.Println("\tUsage:")
		fmt.Println("")
		for _, cmd := range commands {
			fmt.Printf("\t\t%s: %s\n", cmd.name, cmd.description)
		}
		return nil
	}
}

func commandMap(configuration *config) error {
	locations, previous, next := pokedex.GetLocations(configuration.next)
	configuration.previous = previous
	configuration.next = next
	for _, location := range locations {
		fmt.Println("\t", location)
	}
	return nil
}

func commandMapb(configuration *config) error {
	locations, previous, next := pokedex.GetLocations(configuration.previous)
	configuration.previous = previous
	configuration.next = next
	for _, location := range locations {
		fmt.Println("\t", location)
	}
	return nil
}

func exploreMap(configuration *config) error {
	pokemons, available := pokedex.GetPokemons("https://pokeapi.co/api/v2/location-area/" + configuration.variable)
	if !available {
		return nil
	}
	for _, pokemon := range pokemons {
		fmt.Println("\t", pokemon)
	}
	return nil
}

func catchPokemon(configuration *config) error {
	fmt.Println("Throwing a Pokeball at " + configuration.variable + "...")
	caught, err := pokedex.CatchPokemon(configuration.variable)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if caught {
		fmt.Println(configuration.variable + " was caught!")
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Println(configuration.variable + " escaped!")
	}
	return nil
}

func inspectPokemon(configuration *config) error {
	pokedex.Inspect(configuration.variable)
	return nil
}

func viewPokedex(configuration *config) error {
	pokedex.ViewPokedex()
	return nil
}

func main() {

	configuration := config{next: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"}

	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex.",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Paginates forward through maps in pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Paginates backwards through maps in pokemon world.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "explore <map> Explores avaialble Pokémon in provided map.",
			callback:    exploreMap,
		},
		"catch": {
			name:        "catch",
			description: "Throw a pokéball to a pokémon, as possibly catch it.",
			callback:    catchPokemon,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught pokemon.",
			callback:    inspectPokemon,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List caught pokemons.",
			callback:    viewPokedex,
		},
	}
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    makeHelpCommand(commands),
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			return
		}
		err := scanner.Err()
		if err != nil {
			error := fmt.Errorf("Error Scanning %w", err)
			fmt.Println(error)
			return
		}
		data := scanner.Text()
		dataElements := pokedex.CleanInput(data)
		if len(dataElements) == 0 {
			continue
		}
		command := dataElements[0]
		if len(dataElements) > 1 {
			configuration.variable = dataElements[1]
		}
		if executeCommand, exists := commands[command]; exists {
			executeCommand.callback(&configuration)
		} else {
			fmt.Println("Unknown command")
		}
	}
}
