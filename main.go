package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const POKEDEX_PROMPT = "Pokedex > "

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(POKEDEX_PROMPT)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("Your command was:", cleanInput(text)[0])
		fmt.Print(POKEDEX_PROMPT)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
