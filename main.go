package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	pokecache "github.com/AaroLankinen/pokedex_cli/internal"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, []string) error
}

type Config struct {
	cache    *pokecache.Cache
	next     string
	previous string
}

type RespLocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type RespLocationArea struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandMap(cfg *Config, args []string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.next != "" {
		url = cfg.next
	}

	var body []byte
	if val, ok := cfg.cache.Get(url); ok {
		body = val
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		cfg.cache.Add(url, body)
	}

	resStruct := RespLocationAreas{}
	err := json.Unmarshal(body, &resStruct)
	if err != nil {
		return err
	}

	if resStruct.Next != nil {
		cfg.next = *resStruct.Next
	} else {
		cfg.next = ""
	}

	if resStruct.Previous != nil {
		cfg.previous = *resStruct.Previous
	} else {
		cfg.previous = ""
	}

	for _, area := range resStruct.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(cfg *Config, args []string) error {
	if cfg.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	var body []byte
	if val, ok := cfg.cache.Get(cfg.previous); ok {
		body = val
	} else {
		res, err := http.Get(cfg.previous)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		cfg.cache.Add(cfg.previous, body)
	}

	resStruct := RespLocationAreas{}
	err := json.Unmarshal(body, &resStruct)
	if err != nil {
		return err
	}

	if resStruct.Next != nil {
		cfg.next = *resStruct.Next
	} else {
		cfg.next = ""
	}

	if resStruct.Previous != nil {
		cfg.previous = *resStruct.Previous
	} else {
		cfg.previous = ""
	}

	for _, area := range resStruct.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandExplore(cfg *Config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: explore <location_name>")
	}

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", args[0])
	var body []byte
	if val, ok := cfg.cache.Get(url); ok {
		body = val
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		cfg.cache.Add(url, body)
	}

	resStruct := RespLocationArea{}
	err := json.Unmarshal(body, &resStruct)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon in %s:\n", args[0])
	for _, pokemon := range resStruct.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Lists the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location area to find Pokemon",
			callback:    commandExplore,
		},
	}
}

func commandExit(cfg *Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func main() {
	cfg := &Config{
		cache: pokecache.NewCache(5 * time.Minute),
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Pokedex CLI")
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		args := cleanInput(input)
		if len(args) == 0 {
			continue
		}

		commandName := args[0]
		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(cfg, args[1:])
		if err != nil {
			fmt.Println(err)
		}
	}
}
