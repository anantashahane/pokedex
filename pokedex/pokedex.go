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

type PokeLocations struct {
	Count    int             `json:"count"`
	Next     string          `json:"next"`
	Previous string          `json:"previous"`
	Results  []PokemonEntity `json:"results"`
}

type PokemonEntity struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationInfo struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int           `json:"game_index"`
	ID        int           `json:"id"`
	Location  PokemonEntity `json:"location"`
	Name      string        `json:"name"`
	Names     []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon        PokemonEntity `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int           `json:"max_chance"`
			Version   PokemonEntity `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Abilities []struct {
		Ability  PokemonEntity `json:"ability"`
		IsHidden bool          `json:"is_hidden"`
		Slot     int           `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Cries          struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Forms       []PokemonEntity `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	Height    int `json:"height"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"item"`
		VersionDetails []struct {
			Rarity  int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	ID                     int    `json:"id"`
	IsDefault              bool   `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			Order        any `json:"order"`
			VersionGroup struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name          string `json:"name"`
	Order         int    `json:"order"`
	PastAbilities []struct {
		Abilities []struct {
			Ability  any  `json:"ability"`
			IsHidden bool `json:"is_hidden"`
			Slot     int  `json:"slot"`
		} `json:"abilities"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
	} `json:"past_abilities"`
	PastTypes []any `json:"past_types"`
	Species   struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      string `json:"back_default"`
		BackFemale       string `json:"back_female"`
		BackShiny        string `json:"back_shiny"`
		BackShinyFemale  string `json:"back_shiny_female"`
		FrontDefault     string `json:"front_default"`
		FrontFemale      string `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale string `json:"front_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontDefault string `json:"front_default"`
				FrontFemale  any    `json:"front_female"`
			} `json:"dream_world"`
			Home struct {
				FrontDefault     string `json:"front_default"`
				FrontFemale      string `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale string `json:"front_shiny_female"`
			} `json:"home"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"official-artwork"`
			Showdown struct {
				BackDefault      string `json:"back_default"`
				BackFemale       string `json:"back_female"`
				BackShiny        string `json:"back_shiny"`
				BackShinyFemale  any    `json:"back_shiny_female"`
				FrontDefault     string `json:"front_default"`
				FrontFemale      string `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale string `json:"front_shiny_female"`
			} `json:"showdown"`
		} `json:"other"`
		Versions struct {
			GenerationI struct {
				RedBlue struct {
					BackDefault      string `json:"back_default"`
					BackGray         string `json:"back_gray"`
					BackTransparent  string `json:"back_transparent"`
					FrontDefault     string `json:"front_default"`
					FrontGray        string `json:"front_gray"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"red-blue"`
				Yellow struct {
					BackDefault      string `json:"back_default"`
					BackGray         string `json:"back_gray"`
					BackTransparent  string `json:"back_transparent"`
					FrontDefault     string `json:"front_default"`
					FrontGray        string `json:"front_gray"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"yellow"`
			} `json:"generation-i"`
			GenerationIi struct {
				Crystal struct {
					BackDefault           string `json:"back_default"`
					BackShiny             string `json:"back_shiny"`
					BackShinyTransparent  string `json:"back_shiny_transparent"`
					BackTransparent       string `json:"back_transparent"`
					FrontDefault          string `json:"front_default"`
					FrontShiny            string `json:"front_shiny"`
					FrontShinyTransparent string `json:"front_shiny_transparent"`
					FrontTransparent      string `json:"front_transparent"`
				} `json:"crystal"`
				Gold struct {
					BackDefault      string `json:"back_default"`
					BackShiny        string `json:"back_shiny"`
					FrontDefault     string `json:"front_default"`
					FrontShiny       string `json:"front_shiny"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"gold"`
				Silver struct {
					BackDefault      string `json:"back_default"`
					BackShiny        string `json:"back_shiny"`
					FrontDefault     string `json:"front_default"`
					FrontShiny       string `json:"front_shiny"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"silver"`
			} `json:"generation-ii"`
			GenerationIii struct {
				Emerald struct {
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"emerald"`
				FireredLeafgreen struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"firered-leafgreen"`
				RubySapphire struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"ruby-sapphire"`
			} `json:"generation-iii"`
			GenerationIv struct {
				DiamondPearl struct {
					BackDefault      string `json:"back_default"`
					BackFemale       string `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  string `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					BackDefault      string `json:"back_default"`
					BackFemale       string `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  string `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"heartgold-soulsilver"`
				Platinum struct {
					BackDefault      string `json:"back_default"`
					BackFemale       string `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  string `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"platinum"`
			} `json:"generation-iv"`
			GenerationV struct {
				BlackWhite struct {
					Animated struct {
						BackDefault      string `json:"back_default"`
						BackFemale       string `json:"back_female"`
						BackShiny        string `json:"back_shiny"`
						BackShinyFemale  string `json:"back_shiny_female"`
						FrontDefault     string `json:"front_default"`
						FrontFemale      string `json:"front_female"`
						FrontShiny       string `json:"front_shiny"`
						FrontShinyFemale string `json:"front_shiny_female"`
					} `json:"animated"`
					BackDefault      string `json:"back_default"`
					BackFemale       string `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  string `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationVi struct {
				OmegarubyAlphasapphire struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"omegaruby-alphasapphire"`
				XY struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"x-y"`
			} `json:"generation-vi"`
			GenerationVii struct {
				Icons struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  any    `json:"front_female"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"ultra-sun-ultra-moon"`
			} `json:"generation-vii"`
			GenerationViii struct {
				Icons struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  string `json:"front_female"`
				} `json:"icons"`
			} `json:"generation-viii"`
		} `json:"versions"`
	} `json:"sprites"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

var cache = pokecache.NewCache(time.Second * 10)
var caughtPokemon = map[string]Pokemon{}

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
		resp, err := http.Get(url)
		if err != nil {
			return PokeLocations{}, fmt.Errorf("Error getting data from %s: %w", url, err)
		}
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return PokeLocations{}, fmt.Errorf("Error reading body from %s: %w", url, err)
		}
		defer resp.Body.Close()
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
		resp, err := http.Get(url)
		if err != nil {
			return []PokemonEntity{}, fmt.Errorf("Error getting data from %s: %w", url, err)
		}
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return []PokemonEntity{}, fmt.Errorf("Error reading body from %s: %w", url, err)
		}
		defer resp.Body.Close()
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
		resp, err := http.Get(url)
		if err != nil {
			return pokemon, fmt.Errorf("Error getting data from %s: %w", url, err)
		}
		defer resp.Body.Close()
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return pokemon, fmt.Errorf("Error reading body obtained from %s: %w", url, err)
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
