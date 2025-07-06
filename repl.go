package main

import (
	"bufio"
	"fmt"
	"github.com/D3rise/pokedexcli/internal/commands"
	"github.com/D3rise/pokedexcli/internal/context"
	"github.com/D3rise/pokedexcli/internal/pokedex"
	"github.com/D3rise/pokedexcli/pokedexapi"
	"log"
	"os"
	"strings"
)

const (
	PokedexPrompt = "Pokedex > "
	PokedexMotd   = "Welcome to the Pokedex!"
)

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	registry := commands.GetCommandRegistry()

	ctx := context.NewContext()
	initializeContext(ctx)

	fmt.Println(PokedexMotd)
	fmt.Print(PokedexPrompt)
	for scanner.Scan() {
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			fmt.Print(PokedexPrompt)
			continue
		}

		cmd := input[0]
		args := input[1:]

		if c, ok := registry[cmd]; ok && len(args) == len(c.Args) {
			err := c.Callback(ctx, args...)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("Unknown command or wrong number of arguments.")
		}

		fmt.Print(PokedexPrompt)
	}
}

func initializeContext(c *context.Context) {
	c.Set(pokedexapi.PokedexApiContextKey, pokedexapi.NewPokedexAPI(""))
	c.Set(pokedex.PokedexContextKey, pokedex.NewPokedex())
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
