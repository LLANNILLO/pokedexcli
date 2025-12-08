package main

import (
	"errors"
	"fmt"

	"github.com/llannillo/pokedexcli/internal/pokeapi"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name to inspect")
	}

	name := args[0]
	pokemon, ok := cfg.caughtPokemons[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	inspectPokemon(pokemon)
	return nil
}

func inspectPokemon(pokemon pokeapi.Pokemon) {
	fmt.Println("Name: " + pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	stats := pokemon.Stats
	for _, stat := range stats {
		statName := stat.Stat.Name
		statVal := stat.BaseStat
		fmt.Printf(" -%s: %d\n", statName, statVal)
	}
	fmt.Println("Types:")
	types := pokemon.Types
	for _, ptype := range types {
		typeName := ptype.Type.Name
		fmt.Printf(" - %s \n", typeName)
	}
}
