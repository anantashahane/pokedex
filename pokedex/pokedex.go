package pokedex

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"strings"
	"time"

	pokecache "github.com/anantashahane/pokedex/pokecache"
)

var cache = pokecache.NewCache(time.Minute * 2)
var caughtPokemon = map[string]Pokemon{}

func fetchData(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("Error getting data from %s: %w", url, err)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Error reading body from %s: %w", url, err)
	}
	defer resp.Body.Close()
	return data, err
}

func GetLocations(url string) (locations []string, previous, next string) {
	locationsData, err := fetchLocationsData(url)
	if err != nil {
		return []string{}, "", ""
	}
	locations = []string{}

	for _, location := range locationsData.Results {
		locations = append(locations, location.Name)
	}
	return locations, locationsData.Previous, locationsData.Next
}

func fetchLocationsData(url string) (locations PokeLocations, err error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	}
	data, available := cache.Get(url)
	if !available {
		data, err = fetchData(url)
		if err != nil {
			return PokeLocations{}, err
		}
	}
	cache.Add(url, data)

	var locationData PokeLocations
	err = json.Unmarshal(data, &locationData)
	if err != nil {
		return PokeLocations{}, fmt.Errorf("Error decoding data received from %s: %w", url, err)
	}
	return locationData, nil
}

func CleanInput(text string) (word []string) {
	words := strings.Fields(strings.ToLower(text))
	if len(words) == 0 {
		return []string{}
	}
	return words
}

func GetPokemons(url string) (pokemons []string, available bool) {
	pokemonEntities, err := fetchPokemons(url)
	available = true
	if err != nil {
		fmt.Printf("%s", err)
		available = false
		return pokemons, available
	}
	for _, entity := range pokemonEntities {
		pokemons = append(pokemons, entity.Name)
	}
	return pokemons, available
}

func fetchPokemons(url string) (areas []PokemonEntity, err error) {
	data, available := cache.Get(url)
	if !available {
		data, err = fetchData(url)
		if err != nil {
			return []PokemonEntity{}, err
		}
	}
	cache.Add(url, data)
	var locationData LocationInfo
	err = json.Unmarshal(data, &locationData)
	if err != nil {
		return []PokemonEntity{}, fmt.Errorf("Error decoding data received from %s: %w", url, err)
	}
	pokemons := []PokemonEntity{}
	for _, pokemon := range locationData.PokemonEncounters {
		pokemons = append(pokemons, pokemon.Pokemon)
	}
	return pokemons, nil
}

func CatchPokemon(name string) (caught bool, err error) {
	pokemon, err := fetchPokemon(name)
	if err != nil {
		fmt.Println(err)
		return caught, err
	}
	random := rand.IntN(1000)
	if random > pokemon.BaseExperience {
		caught = true
		caughtPokemon[pokemon.Name] = pokemon
	}
	return caught, nil
}

func fetchPokemon(name string) (pokemon Pokemon, err error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + name
	data, available := cache.Get(url)
	if !available {
		data, err = fetchData(url)
		if err != nil {
			return pokemon, err
		}
	}
	cache.Add(url, data)

	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return pokemon, fmt.Errorf("Error unmarshalling body from %s: %w", url, err)
	}
	return pokemon, nil
}

func Inspect(name string) {
	if pokemon, caught := caughtPokemon[name]; caught {
		fmt.Println("\tName:", pokemon.Name)
		fmt.Println("\tHeight:", pokemon.Height)
		fmt.Println("\tWeight:", pokemon.Weight)
		fmt.Println("\tStats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("\t\t-%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("\tTypes:")
		for _, poketype := range pokemon.Types {
			fmt.Printf("\t\t-%s\n", poketype.Type.Name)
		}
		return
	}
	fmt.Println("\tyou have not caught that pokemon")
}

func ViewPokedex() {
	fmt.Println("\tYour Pokedex:")
	for key, _ := range caughtPokemon {
		fmt.Println("\t\t-", key)
	}
}
