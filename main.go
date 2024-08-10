package main

import (
	"time"

	"github.com/Ace0314/Pokedex-with-Go/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapi: pokeClient,
	}

	startRepl(cfg)
}
