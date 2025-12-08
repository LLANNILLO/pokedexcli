package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name to attempt to catch")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	caught := attemptCatch(pokemon.BaseExperience)
	if caught {
		cfg.caughtPokemons[pokemon.Name] = pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	return nil
}

func attemptCatch(baseExperience int) bool {
	// Calculate capture rate (inversly proportional to experience)
	// Weak Pokemon: high rate; Strong Pokemon: low rate
	catchRate := max(5, 255-baseExperience)

	// Throwing 4 times the rate (shakes of Pokemon)
	for i := 0; i < 4; i++ {
		roll := rand.Intn(255)
		if roll > catchRate {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
