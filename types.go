package main

import "pokedox/internal/pokeapi"

type cliCommand struct {
	name        string
	description string
	execute     func(args ...string) error
}

type Config struct {
	Previous string
	Next     string
	History  []string
	Pokedox  map[string]pokeapi.Pokemon
}

func NewConfig() *Config {
	return &Config{
		Previous: "",
		Next:     "",
		History:  []string{},
		Pokedox:  make(map[string]pokeapi.Pokemon),
	}
}
