# Pokedex CLI

Pokedex CLI is an interactive command-line Pokedex written in Go. It lets you explore the Pokemon world's location areas, discover which Pokemon live in each one, try to catch them, and inspect the ones you've caught. All data comes from the [PokeAPI](https://pokeapi.co/).

## Requirements

Before running Pokedex CLI you'll need the following:

- **Go** (1.25 or later) — needed to build and install the CLI.
- **An internet connection** — Pokemon and location data is fetched live from the PokeAPI.

## Installing

Install the `pokedexcli` CLI with `go install`:

```bash
go install github.com/llannillo/pokedexcli@latest
```

This builds the binary and places it in your `$GOPATH/bin` (usually `~/go/bin`), so make sure that directory is on your `$PATH`.

## Configuration

No configuration is needed. The CLI works out of the box with sensible defaults:

- **HTTP timeout** — requests to the PokeAPI time out after 5 seconds.
- **Response cache** — API responses are cached in memory and expire after 25 minutes, so revisiting a page or Pokemon is instant.

Your caught Pokemon live in memory only, so your Pokedex is reset every time you exit the CLI.

## Running

Start the interactive REPL:

```bash
pokedexcli
```

Then type commands at the `Pokedex >` prompt:

```text
Pokedex > map
Pokedex > explore pastoria-city-area
Pokedex > catch pikachu
Pokedex > inspect pikachu
Pokedex > pokedex
```

The commands available:

| Command | Description |
| --- | --- |
| `help` | Display a help message with all commands. |
| `map` | Show the next page of location areas. |
| `mapb` | Show the previous page of location areas. |
| `explore <area_name>` | List the Pokemon found in a location area. |
| `catch <pokemon_name>` | Throw a Pokeball and try to catch a Pokemon. Stronger Pokemon are harder to catch. |
| `inspect <pokemon_name>` | Show the height, weight, stats, and types of a Pokemon you've caught. |
| `pokedex` | List all the Pokemon you've caught. |
| `exit` | Exit the Pokedex. |

## Testing

Run the test suite with:

```bash
go test ./...
```
