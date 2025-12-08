package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for pokemonName := range cfg.caughtPokemons {
		fmt.Printf(" - %s\n", pokemonName)
	}
	return nil
}
