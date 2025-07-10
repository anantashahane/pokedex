package main

import (
	"fmt"
	"os"

	pokedex "github.com/anantashahane/pokedex/pokedex"
)

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
