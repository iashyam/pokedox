package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"pokedox/internal/pokeapi"
	"pokedox/internal/pokecache"
	"time"
)

func getCommands(conf *Config, cache *pokecache.PokeCache) map[string]cliCommand {

	// this map will abstract the commands available in the
	commands := map[string]cliCommand{
		"exit": cliCommand{
			name:        "exit",
			description: "Exit the Pokedox application",
			execute:     commandExit,
		},
		"quit": cliCommand{
			name:        "exit",
			description: "Exit the Pokedox application",
			execute:     commandExit,
		},
		"history": cliCommand{
			name:        "history",
			description: "Prints the list of commands run in the session",
			execute:     func(args ...string) error { return commandHistory(conf, args...) },
		},
		"help": cliCommand{
			name:        "help",
			description: "Display this help message",
			execute:     commandHelp,
		},
		"echo": cliCommand{
			name:        "echo {argument}",
			description: "Prints the argument on the terminal.",
			execute:     commandEcho,
		},
		"map": cliCommand{
			name:        "map",
			description: "Display a list of next location areas",
			execute:     func(args ...string) error { return commandMap(conf, cache, args...) },
		},
		"mapb": cliCommand{
			name:        "map",
			description: "Display a list of previous location areas",
			execute:     func(args ...string) error { return commandBMap(conf, cache, args...) },
		},
		"explore": cliCommand{
			name:        "explore {location}",
			description: "Display a list of Pokemon in a location area",
			execute:     func(args ...string) error { return commandExplore(cache, args...) },
		},
		"pokedox": cliCommand{
			name:        "pokedox",
			description: "Lists all of the pokemon you have caught!",
			execute:     func(args ...string) error { return commandList(conf) },
		}, "catch": cliCommand{
			name:        "catch {pokemon name}",
			description: "Throws a ball to catch the pokemon given",
			execute:     func(args ...string) error { return commandCatch(conf, cache, args...) },
		},
		"inspect": cliCommand{
			name:        "inspect {pokemon name}",
			description: "Shows some interesting stats about the pokemon",
			execute:     func(args ...string) error { return commandInspect(conf, args...) },
		},
	}

	return commands
}

/// Definge all the commands below and add it to commands struct

func commandEcho(args ...string) error {
	if len(args) <= 0 {
		return fmt.Errorf("no argument passed")
	}
	for _, arg := range args {
		fmt.Println(arg)
	}
	return nil
}

func commandHistory(conf *Config, args ...string) error {
	if len(conf.History) == 0 {
		fmt.Println("No commands in history.")
		return nil
	}
	for i, cmd := range conf.History {
		fmt.Printf(" %d: %s \n", i+1, cmd)
	}
	return nil
}

func commandExit(args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(args ...string) error {
	commands := getCommands(conf, cache)
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for name, command := range commands {
		fmt.Printf("%s: %s\n", name, command.description)
	}
	return nil
}

func commandMap(ptr *Config, cache *pokecache.PokeCache, args ...string) error {
	url := ptr.Next

	if ptr.Next == "" && ptr.Previous == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	if ptr.Next == "" && ptr.Previous != "" {
		fmt.Println("You are at the last page!")
		return nil
	}

	var locationResponse []byte
	var err error
	if val, ok := cache.Get(url); ok {
		// fmt.Println("cache hit")
		// fmt.Println("----------")
		locationResponse = val
	} else {
		locationResponse, err = pokeapi.GetLocationResponse(url)
		// fmt.Println("Cache miss")
		// fmt.Println("----------")

		if err != nil {
			return err
		}
		cache.Add(url, locationResponse)
	}

	location_list, err := pokeapi.ParseLocationResponse(locationResponse)
	if err != nil {
		return err
	}

	for _, location := range location_list.Results {
		fmt.Println("- ", location["name"])
	}
	ptr.Next = location_list.Next
	ptr.Previous = location_list.Previous

	return nil
}

func commandBMap(ptr *Config, cache *pokecache.PokeCache, args ...string) error {
	url := ptr.Previous

	if ptr.Next == "" && ptr.Previous == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	if ptr.Next != "" && ptr.Previous == "" {
		fmt.Println("You are at the first page!")
		return nil
	}

	var locationResponse []byte
	var err error
	val, ok := cache.Get(url)
	if ok {
		fmt.Println("cache hit")
		fmt.Println("----------")
		locationResponse = val
	} else {
		fmt.Println("cache miss")
		fmt.Println("----------")
		locationResponse, err = pokeapi.GetLocationResponse(url)
		if err != nil {
			return err
		}
		cache.Add(url, locationResponse)
	}

	location_list, err := pokeapi.ParseLocationResponse(locationResponse)
	if err != nil {
		return err
	}

	for _, location := range location_list.Results {
		fmt.Println("- ", location["name"])
	}
	ptr.Next = location_list.Next
	ptr.Previous = location_list.Previous

	return nil
}

func commandExplore(cache *pokecache.PokeCache, args ...string) error {
	if len(args) <= 0 {
		return fmt.Errorf("no location area provided")
	}
	location := args[0]
	url := "https://pokeapi.co/api/v2/location-area/" + location

	var encounterResponse []byte
	var err error
	if val, ok := cache.Get(url); ok {
		encounterResponse = val
	} else {
		encounterResponse, err = pokeapi.GetLocationAreaEncounterResponse(url)
		if err != nil {
			return err
		}
		cache.Add(url, encounterResponse)
	}

	encounterResults, err := pokeapi.ParseLocationAreaEncounterResponse(encounterResponse)
	if err != nil {
		return err
	}

	fmt.Printf("Pokémon in %s:\n", location)
	for _, encounter := range encounterResults.PokemonEncounters {
		fmt.Println("- ", encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(conf *Config, cache *pokecache.PokeCache, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you can catch one and only one pokemon")
	}

	pokemon := args[0]
	if _, ok := conf.Pokedox[pokemon]; ok {
		fmt.Println("You already have, ", pokemon)
		return nil
	}
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon

	var pokemonResponse []byte
	var err error
	if val, ok := cache.Get(url); ok {
		pokemonResponse = val
	} else {
		pokemonResponse, err = pokeapi.GetPokemonResponse(url)
		if err != nil {
			return err
		}
		cache.Add(url, pokemonResponse)
	}

	pokemonResults, err := pokeapi.ParsePokemonResponse(pokemonResponse)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s", pokemon)
	time.Sleep(time.Second)
	fmt.Println("....")

	baseExperience := float64(pokemonResults.BaseExperience) / 1000.0
	rand := rand.Float64()
	fmt.Printf("experience %f, rand %f\n", baseExperience, rand)
	if rand > baseExperience {
		fmt.Println("Caught", pokemon)
		conf.Pokedox[pokemon] = pokemonResults
		return nil
	}
	fmt.Printf("%s escaped\n", pokemon)
	return nil
}

func commandList(conf *Config) error {
	list := conf.Pokedox
	if len(list) == 0 {
		fmt.Println("You haven't caught any pokemon yet.")
		return nil
	}
	fmt.Println("Your Pokedox:")
	for _, pokemon := range list {
		fmt.Println("- ", pokemon.Name)
	}

	return nil
}

func commandInspect(conf *Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you can only inspect one and only one pokemon")
	}
	if len(args) == 0 {
		return fmt.Errorf("enter the name of pokemon form you pokedox to inspect")
	}

	pokemonName := args[0]
	pokemon, ok := conf.Pokedox[pokemonName]
	if !ok {
		return fmt.Errorf("you haven't caught %s yet", pokemonName)
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s \n", stat.Stat.Name)
	}
	fmt.Println("Types:")
	for _, type_ := range pokemon.Types {
		fmt.Printf("- %s \n", type_.Type.Name)
	}
	return nil
}
