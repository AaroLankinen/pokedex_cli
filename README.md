# Pokedex CLI

A command-line Pokedex application written in Go. This project interacts with the [PokéAPI](https://pokeapi.co/) to provide information about Pokémon locations, species, and stats.

## Features

- **Caching**: Local caching mechanism to reduce API calls and improve performance.
- **Exploration**: Browse location areas and see which Pokémon can be found there.
- **Catching**: Try your luck catching Pokémon! Success rate depends on the Pokémon's base experience.
- **Pokedex**: Manage your collection of caught Pokémon.
- **Inspection**: View detailed stats and types for any Pokémon in your Pokedex.

## Installation

Ensure you have [Go](https://go.dev/doc/install) installed.

1. Clone the repository.
2. Navigate to the project directory.
3. Build the application:
   ```bash
   go build .
   ```

## Usage

Run the executable to start the interactive REPL:

```bash
./pokedex_cli
```

### Available Commands

- `help`: Displays a help message with all available commands.
- `map`: Lists the next 20 location areas in the Pokémon world.
- `mapb`: Lists the previous 20 location areas.
- `explore <location_name>`: Lists all Pokémon encountered in a specific area.
- `catch <pokemon_name>`: Attempts to catch a Pokémon and add it to your Pokedex.
- `inspect <pokemon_name>`: Shows stats, height, weight, and types of a Pokémon you have caught.
- `pokedex`: Prints a list of all the Pokémon you have caught.
- `exit`: Closes the application.