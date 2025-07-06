package main

import (
	"bufio"
	"fmt"
	"github.com/D3rise/pokedexcli/internal/commands"
	"github.com/D3rise/pokedexcli/internal/context"
	"github.com/D3rise/pokedexcli/pokedexapi"
	"log"
	"os"
	"strings"
)

const (
	POKEDEX_PROMPT = "Pokedex > "
	POKEDEX_MOTD   = "Welcome to the Pokedex!"
)

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	registry := commands.GetCommandRegistry()

	ctx := context.NewContext()
	initializeContext(ctx)

	fmt.Println(POKEDEX_MOTD)
	fmt.Print(POKEDEX_PROMPT)
	for scanner.Scan() {
		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			fmt.Print(POKEDEX_PROMPT)
			continue
		}

		if c, ok := registry[text[0]]; ok {
			err := c.Callback(ctx, text[1:]...)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
		fmt.Print(POKEDEX_PROMPT)
	}
}

func initializeContext(c *context.Context) {
	c.Set(context.PokedexAPI, pokedexapi.NewPokedexAPI(""))
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
